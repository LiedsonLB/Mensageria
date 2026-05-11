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
	log.Println("🏛️ Worker SEFAZ: conectando à fila...")
	
	repo := repository.NewPedidoRepository(db)

	msgs, err := rabbit.Consume("nota_fiscal.pronta")
	if err != nil {
		log.Printf("❌ Erro ao consumir fila nota_fiscal.pronta: %v", err)
		return
	}

	log.Println("✅ Worker SEFAZ aguardando notas para autorizacao...")

	for msg := range msgs {
		log.Println("📨 Mensagem recebida na fila nota_fiscal.pronta!")
		
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

		log.Printf("🏛️ Enviando nota fiscal para SEFAZ: %s", pedido.ID)

		repo.AddStatusHistory(pedido.ID, "ENVIANDO_SEFAZ", "Enviando nota para SEFAZ")

		xmlPath := fmt.Sprintf("notas-fiscais/%s.xml", pedido.ID)
		resposta, err := sefaz.EnviarNFe(xmlPath, pedido)
		if err != nil {
			log.Printf("❌ Erro ao enviar para SEFAZ: %v", err)
			repo.AddStatusHistory(pedido.ID, "ERRO_SEFAZ", err.Error())
			continue
		}

		log.Printf("✅ Nota autorizada pela SEFAZ: %s - Protocolo: %s", pedido.ID, resposta.Protocolo)

		repo.AddStatusHistory(pedido.ID, "NF_AUTORIZADA", 
			fmt.Sprintf("Nota autorizada pela SEFAZ. Protocolo: %s", resposta.Protocolo))

		eventoPDF := map[string]interface{}{
			"evento":           "nfe.autorizada",
			"pedido_id":        pedido.ID,
			"pedido":           pedido,
			"protocolo":        resposta.Protocolo,
			"chave_acesso":     resposta.ChaveAcesso,
			"data_autorizacao": resposta.DataAutorizacao,
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
			log.Printf("📄 Evento nfe.autorizada publicado para PDF: %s", pedido.ID)
		}
	}
}