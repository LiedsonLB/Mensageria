// internal/worker/pagamento_worker.go
package worker

import (
	"encoding/json"
	"log"
	"time"

	"mensageria-go/internal/model"
	"mensageria-go/internal/queue"
	"mensageria-go/internal/repository"
	"mensageria-go/internal/service"

	"gorm.io/gorm"
)

func StartPagamentoWorker(
	rabbit *queue.RabbitMQ,
	stripeService *service.StripeService,
	db *gorm.DB,
) {
	log.Println("🔄 Worker de Pagamento: conectando à fila...")
	
	repo := repository.NewPedidoRepository(db)

	msgs, err := rabbit.Consume("pagamentos")
	if err != nil {
		log.Printf("❌ Erro ao consumir fila pagamentos: %v", err)
		return
	}

	log.Println("✅ Worker de Pagamento aguardando pedidos...")

	for msg := range msgs {
		var pedido model.Pedido
		err := json.Unmarshal(msg.Body, &pedido)

		if err != nil {
			log.Printf("❌ Erro ao parsear JSON: %v", err)
			continue
		}

		log.Printf("💰 Processando pagamento do pedido: %s - Valor: R$ %.2f", pedido.ID, pedido.ValorTotal)

		repo.UpdateStatus(pedido.ID, "PROCESSANDO_PAGAMENTO")
		repo.AddStatusHistory(pedido.ID, "PROCESSANDO_PAGAMENTO", "Iniciando processamento do pagamento")

		err = stripeService.Pagar(pedido.ValorTotal)
		if err != nil {
			log.Printf("❌ Erro no pagamento do pedido %s: %v", pedido.ID, err)
			repo.UpdateStatus(pedido.ID, "PAGAMENTO_FALHOU")
			repo.AddStatusHistory(pedido.ID, "PAGAMENTO_FALHOU", err.Error())
			continue
		}

		log.Printf("✅ Pagamento aprovado para pedido: %s", pedido.ID)
		
		repo.UpdateStatus(pedido.ID, "PAGAMENTO_APROVADO")
		repo.AddStatusHistory(pedido.ID, "PAGAMENTO_APROVADO", "Pagamento aprovado com sucesso")

		evento := map[string]interface{}{
			"evento":    "pagamento_aprovado",
			"pedido_id": pedido.ID,
			"pedido":    pedido,
			"timestamp": time.Now(),
		}
		
		body, err := json.Marshal(evento)
		if err != nil {
			log.Printf("⚠️ Erro ao serializar evento: %v", err)
			continue
		}

		err = rabbit.Publish("notas_fiscais", body)
		if err != nil {
			log.Printf("⚠️ Erro ao publicar evento para nota fiscal: %v", err)
		} else {
			log.Printf("📤 Evento publicado para notas_fiscais: %s", pedido.ID)
		}
	}
}