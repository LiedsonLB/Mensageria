package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"mensageria-go/internal/model"
	"mensageria-go/internal/queue"
	"mensageria-go/internal/repository"
	"mensageria-go/internal/sefaz"

	"gorm.io/gorm"
)

func StartSefazWorker(
	rabbit *queue.RabbitMQ,
	db *gorm.DB,
) {
	log.Println("🏛️ ========== WORKER SEFAZ ==========")
	log.Println("🔄 Conectando à fila nota_fiscal.pronta...")
	
	repo := repository.NewPedidoRepository(db)

	msgs, err := rabbit.Consume("nota_fiscal.pronta")
	if err != nil {
		log.Printf("❌ Erro ao consumir fila: %v", err)
		return
	}

	log.Println("✅ Worker SEFAZ aguardando notas para autorização...")
	log.Println("📡 Protocolo: SOAP | WebService | Lote ASSÍNCRONO")

	for msg := range msgs {
		log.Println("\n📨 ========== NOVA NOTA RECEBIDA ==========")
		
		var evento map[string]interface{}
		err := json.Unmarshal(msg.Body, &evento)
		if err != nil {
			log.Printf("❌ Erro ao parsear evento: %v", err)
			continue
		}

		pedidoData, err := json.Marshal(evento["pedido"])
		if err != nil {
			log.Printf("❌ Erro ao extrair pedido: %v", err)
			continue
		}

		var pedido model.Pedido
		err = json.Unmarshal(pedidoData, &pedido)
		if err != nil {
			log.Printf("❌ Erro ao parsear pedido: %v", err)
			continue
		}

		log.Printf("📄 Pedido ID: %s", pedido.ID)
		log.Printf("💰 Valor: R$ %.2f", pedido.ValorTotal)

		repo.AddStatusHistory(pedido.ID, "ENVIANDO_SEFAZ", "📤 Enviando nota para SEFAZ via SOAP")

		xmlPath := fmt.Sprintf("notas-fiscais/%s.xml", pedido.ID)
		
		resposta, err := sefaz.EnviarNFe(xmlPath, pedido)
		
		if err != nil {
			log.Printf("❌ Erro ao enviar para SEFAZ: %v", err)
			repo.AddStatusHistory(pedido.ID, "ERRO_SEFAZ", err.Error())
			continue
		}

		log.Printf("\n✅ ========== NOTA AUTORIZADA PELA SEFAZ ==========")
		log.Printf("📋 Status: %s", resposta.Status)
		log.Printf("🔑 Chave de Acesso: %s", resposta.ChaveAcesso)
		log.Printf("📄 Protocolo: %s", resposta.Protocolo)
		log.Printf("📅 Data Autorização: %s", resposta.DataAutorizacao)
		log.Printf("📦 Número do Lote: %s", resposta.NumeroLote)
		log.Printf("📋 Recibo: %s", resposta.Recibo)

		repo.AddStatusHistory(pedido.ID, "NF_AUTORIZADA", 
			fmt.Sprintf("✅ Nota autorizada pela SEFAZ | Protocolo: %s | Chave: %s", 
				resposta.Protocolo, resposta.ChaveAcesso[:20]))

		eventoPDF := map[string]interface{}{
			"evento":           "nfe.autorizada",
			"pedido_id":        pedido.ID,
			"pedido":           pedido,
			"protocolo":        resposta.Protocolo,
			"chave_acesso":     resposta.ChaveAcesso,
			"data_autorizacao": resposta.DataAutorizacao,
			"recibo":           resposta.Recibo,
			"numero_lote":      resposta.NumeroLote,
			"timestamp":        time.Now(),
		}
		
		body, err := json.Marshal(eventoPDF)
		if err != nil {
			log.Printf("⚠️ Erro ao serializar evento: %v", err)
			continue
		}

		err = rabbit.Publish("nfe.autorizada", body)
		if err != nil {
			log.Printf("⚠️ Erro ao publicar evento para PDF: %v", err)
		} else {
			log.Printf("📤 Evento nfe.autorizada publicado para PDF Worker")
		}
	}
}