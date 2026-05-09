package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"mensageria-go/internal/model"
	"mensageria-go/internal/queue"
	"mensageria-go/internal/service"
)

func StartEmailWorker(
	rabbit *queue.RabbitMQ,
	emailService *service.EmailService,
) {
	log.Println("📧 Worker de Email: conectando à fila...")

	msgs, err := rabbit.Consume("emails")
	if err != nil {
		log.Printf("❌ Erro ao consumir fila de emails: %v", err)
		return
	}

	log.Println("✅ Worker de Email aguardando mensagens...")

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

		notaURL, ok := evento["nota_url"].(string)
		if !ok {
			log.Printf("❌ Nota URL não encontrada no evento")
			continue
		}

		notaData, _ := evento["nota_data"].(string)
		if notaData == "" {
			notaData = time.Now().Format("02/01/2006 15:04")
		}

		log.Printf("📧 Preparando email para pedido: %s - Cliente: %s", pedido.ID, pedido.Cliente)

		emailData := model.EmailData{
			PedidoID:   pedido.ID,
			Cliente:    pedido.Cliente,
			Produtos:   pedido.Produtos,
			ValorTotal: fmt.Sprintf("%.2f", pedido.ValorTotal),
			Data:       notaData,
			NotaURL:    notaURL,
		}

		err = emailService.EnviarEmail(pedido.Email, emailData)
		if err != nil {
			log.Printf("❌ Erro ao enviar email para %s: %v", pedido.Email, err)
			continue
		}

		log.Printf("✅ Email enviado com sucesso para: %s (Pedido: %s)", pedido.Email, pedido.ID)
	}
}