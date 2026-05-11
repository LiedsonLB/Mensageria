// internal/worker/pdf_worker.go
package worker

import (
	"encoding/json"
	"log"
	"time"

	"mensageria-go/internal/danfe"
	"mensageria-go/internal/queue"
	"mensageria-go/internal/repository"

	"gorm.io/gorm"
)

func StartPDFWorker(
	rabbit *queue.RabbitMQ,
	db *gorm.DB,
) {
	log.Println("📄 Worker PDF: conectando à fila...")
	
	repo := repository.NewPedidoRepository(db)

	// Consumir eventos de nota autorizada
	msgs, err := rabbit.Consume("nfe.autorizada")
	if err != nil {
		log.Printf("❌ Erro ao consumir fila nfe.autorizada: %v", err)
		return
	}

	log.Println("✅ Worker PDF aguardando notas autorizadas...")

	for msg := range msgs {
		var evento map[string]interface{}
		err := json.Unmarshal(msg.Body, &evento)
		if err != nil {
			log.Printf("❌ Erro ao parsear evento: %v", err)
			continue
		}

		pedidoID, ok := evento["pedido_id"].(string)
		if !ok {
			log.Printf("❌ Pedido ID não encontrado")
			continue
		}

		protocolo, _ := evento["protocolo"].(string)
		chaveAcesso, _ := evento["chave_acesso"].(string)
		
		// Buscar pedido completo
		pedido, err := repo.FindByID(pedidoID)
		if err != nil {
			log.Printf("❌ Erro ao buscar pedido: %v", err)
			continue
		}

		log.Printf("📄 Gerando DANFE para pedido: %s", pedidoID)

		// repo.UpdateStatus(pedidoID, "GERANDO_DANFE")
		repo.AddStatusHistory(pedidoID, "GERANDO_DANFE", "Gerando DANFE da nota fiscal")

		// Gerar DANFE (agora salva em notas-fiscais-pdf/)
		pdfPath, err := danfe.GerarDANFE(pedido, protocolo, chaveAcesso)
		if err != nil {
			log.Printf("❌ Erro ao gerar DANFE: %v", err)
			repo.AddStatusHistory(pedidoID, "ERRO_DANFE", err.Error())
			continue
		}

		log.Printf("✅ DANFE gerado: %s", pdfPath)

		repo.AddStatusHistory(pedidoID, "DANFE_GERADO", "DANFE gerado com sucesso")

		pdfURL := "http://localhost:8080/notas-fiscais-pdf/" + pedidoID + "-danfe.pdf"

		// Publicar evento para email com PDF anexado
		eventoEmail := map[string]interface{}{
			"evento":        "danfe.gerado",
			"pedido_id":     pedidoID,
			"pedido":        pedido,
			"pdf_path":      pdfPath,
			"pdf_url":       pdfURL,
			"protocolo":     protocolo,
			"chave_acesso":  chaveAcesso,
			"timestamp":     time.Now(),
		}
		
		body, err := json.Marshal(eventoEmail)
		if err != nil {
			log.Printf("⚠️ Erro ao serializar evento: %v", err)
			continue
		}

		err = rabbit.Publish("danfe.gerado", body)
		if err != nil {
			log.Printf("⚠️ Erro ao publicar evento para email: %v", err)
		} else {
			log.Printf("📧 Evento danfe.gerado publicado: %s", pedidoID)
		}
	}
}