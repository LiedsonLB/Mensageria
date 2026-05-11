// internal/worker/email_worker.go - CORRIGIDO
package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"os"
	"encoding/base64"

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

	// 🔥 CORRIGIDO: Consumir da fila danfe.gerado
	msgs, err := rabbit.Consume("danfe.gerado")
	if err != nil {
		log.Printf("❌ Erro ao consumir fila danfe.gerado: %v", err)
		return
	}

	log.Println("✅ Worker de Email aguardando DANFES prontos...")

	for msg := range msgs {
		log.Println("📨 Mensagem recebida na fila danfe.gerado!")

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

		log.Printf("📧 Processando pedido: %s", pedidoID)

		// Buscar pedido completo do banco
		pedidoCompleto, err := repo.FindByID(pedidoID)
		if err != nil {
			log.Printf("❌ Erro ao buscar pedido: %v", err)
			continue
		}

		// Pegar a URL do PDF do evento
		pdfURL, ok := evento["pdf_url"].(string)
		if !ok {
			log.Printf("❌ PDF URL não encontrada no evento")
			continue
		}

		log.Printf("📧 Preparando email para pedido: %s - Cliente: %s", pedidoCompleto.ID, pedidoCompleto.Cliente)
		log.Printf("📦 Produtos: %d itens", len(pedidoCompleto.ProdutosList))
		log.Printf("📎 PDF URL: %s", pdfURL)

		repo.UpdateStatus(pedidoCompleto.ID, "ENVIANDO_EMAIL")
		repo.AddStatusHistory(pedidoCompleto.ID, "ENVIANDO_EMAIL", "Preparando envio de email com DANFE")

		// Preparar produtos
		produtos := pedidoCompleto.ProdutosList
		for i := range produtos {
			produtos[i].Subtotal = produtos[i].Preco * float64(produtos[i].Quantidade)
		}

		// URLs
		xmlURL := fmt.Sprintf("http://localhost:8080/notas-fiscais/%s.xml", pedidoID)
		danfeURL := fmt.Sprintf("http://localhost:8080/notas-fiscais-pdf/%s-danfe.pdf", pedidoID)

		// Chave de acesso e protocolo do evento
		chaveAcesso, _ := evento["chave_acesso"].(string)
		protocolo, _ := evento["protocolo"].(string)

		logoBase64 := ""
		logoPath := "assets/neoshop_logo.png"
		if logoData, err := os.ReadFile(logoPath); err == nil {
			logoBase64 = base64.StdEncoding.EncodeToString(logoData)
			log.Println("✅ Logo carregada para o email")
		} else {
			log.Printf("⚠️ Logo não encontrada: %s", logoPath)
		}

		emailData := model.EmailData{
			PedidoID:    pedidoCompleto.ID,
			Cliente:     pedidoCompleto.Cliente,
			Produtos:    produtos,
			ValorTotal:  fmt.Sprintf("%.2f", pedidoCompleto.ValorTotal),
			Data:        time.Now().Format("02/01/2006 15:04"),
			NotaURL:     xmlURL,
			DANFE_URL:   danfeURL,
			ChaveAcesso: chaveAcesso,
			Protocolo:   protocolo,
			Ano:         time.Now().Year(),
			LogoBase64:  logoBase64,
		}

		log.Printf("📧 Enviando email para: %s", pedidoCompleto.Email)

		err = emailService.EnviarEmail(pedidoCompleto.Email, emailData)
		if err != nil {
			log.Printf("❌ Erro ao enviar email para %s: %v", pedidoCompleto.Email, err)
			repo.AddStatusHistory(pedidoCompleto.ID, "EMAIL_FALHOU", err.Error())
			continue
		}

		repo.UpdateStatus(pedidoCompleto.ID, "CONCLUIDO")
		repo.AddStatusHistory(pedidoCompleto.ID, "CONCLUIDO", "Pedido concluído com sucesso - Email com DANFE enviado")

		log.Printf("✅ Email com DANFE enviado com sucesso para: %s (Pedido: %s)", pedidoCompleto.Email, pedidoCompleto.ID)
	}
}
