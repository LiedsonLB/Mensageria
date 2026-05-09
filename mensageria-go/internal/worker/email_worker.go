// internal/worker/email_worker.go
package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"mensageria-go/internal/model"
	"mensageria-go/internal/queue"
	"mensageria-go/internal/repository"
	"mensageria-go/internal/service"

	"gorm.io/gorm"
)

func StartEmailWorker(
	rabbit *queue.RabbitMQ,
	emailService *service.EmailService,
	db *gorm.DB,
) {
	log.Println("📧 Worker de Email: conectando à fila...")
	
	repo := repository.NewPedidoRepository(db)

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

		// Buscar pedido completo do banco para garantir os produtos
		pedidoCompleto, err := repo.FindByID(pedido.ID)
		if err != nil {
			log.Printf("❌ Erro ao buscar pedido completo: %v", err)
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

		log.Printf("📧 Preparando email para pedido: %s - Cliente: %s", pedidoCompleto.ID, pedidoCompleto.Cliente)
		log.Printf("📦 Produtos: %d itens", len(pedidoCompleto.ProdutosList))

		repo.UpdateStatus(pedidoCompleto.ID, "ENVIANDO_EMAIL")
		repo.AddStatusHistory(pedidoCompleto.ID, "ENVIANDO_EMAIL", "Preparando envio de email")

		// Calcular subtotais se necessário
		produtos := pedidoCompleto.ProdutosList
		for i := range produtos {
			produtos[i].Subtotal = produtos[i].Preco * float64(produtos[i].Quantidade)
		}

		log.Printf("📧 Enviando email para: %s", pedidoCompleto.Email)

		emailData := model.EmailData{
			PedidoID:   pedidoCompleto.ID,
			Cliente:    pedidoCompleto.Cliente,
			Produtos:   produtos,
			ValorTotal: fmt.Sprintf("%.2f", pedidoCompleto.ValorTotal),
			Data:       notaData,
			NotaURL:    notaURL,
		}

		err = emailService.EnviarEmail(pedidoCompleto.Email, emailData)
		if err != nil {
			log.Printf("❌ Erro ao enviar email para %s: %v", pedidoCompleto.Email, err)
			repo.AddStatusHistory(pedidoCompleto.ID, "EMAIL_FALHOU", err.Error())
			continue
		}

		repo.UpdateStatus(pedidoCompleto.ID, "CONCLUIDO")
		repo.AddStatusHistory(pedidoCompleto.ID, "CONCLUIDO", "Pedido concluído com sucesso - Email enviado")

		log.Printf("✅ Email enviado com sucesso para: %s (Pedido: %s)", pedidoCompleto.Email, pedidoCompleto.ID)
	}
}