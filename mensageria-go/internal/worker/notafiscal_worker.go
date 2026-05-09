package worker

import (
	"encoding/json"
	"log"
	"time"

	"mensageria-go/internal/model"
	"mensageria-go/internal/queue"
	"mensageria-go/internal/service"
)

func StartNotaFiscalWorker(
	rabbit *queue.RabbitMQ,
) {
	log.Println("📄 Worker de Nota Fiscal: conectando à fila...")

	msgs, err := rabbit.Consume("notas_fiscais")
	if err != nil {
		log.Printf("❌ Erro ao consumir fila de notas fiscais: %v", err)
		return
	}

	log.Println("✅ Worker de Nota Fiscal aguardando eventos...")

	for msg := range msgs {
		var evento map[string]interface{}
		err := json.Unmarshal(msg.Body, &evento)
		if err != nil {
			log.Printf("❌ Erro ao parsear evento: %v", err)
			continue
		}

		// Extrair pedido do evento
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

		log.Printf("📄 Gerando nota fiscal para pedido: %s", pedido.ID)

		// Gerar nota fiscal
		nota, err := service.GerarNota(pedido)
		if err != nil {
			log.Printf("❌ Erro ao gerar nota fiscal: %v", err)
			continue
		}

		notaURL := "http://localhost:8080/" + nota.Arquivo
		log.Printf("✅ Nota fiscal gerada: %s", nota.Arquivo)

		// Publicar evento para o worker de email
		eventoEmail := map[string]interface{}{
			"evento":    "nota_fiscal_gerada",
			"pedido":    pedido,
			"nota_url":  notaURL,
			"nota_data": nota.Data.Format("02/01/2006 15:04"),
			"timestamp": time.Now(),
		}
		
		body, err := json.Marshal(eventoEmail)
		if err != nil {
			log.Printf("⚠️ Erro ao serializar evento de email: %v", err)
			continue
		}

		err = rabbit.Publish("emails", body)
		if err != nil {
			log.Printf("⚠️ Erro ao publicar evento para email: %v", err)
		} else {
			log.Printf("📧 Evento publicado para emails: %s", pedido.ID)
		}
	}
}