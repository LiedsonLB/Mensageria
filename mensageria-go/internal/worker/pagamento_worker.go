package worker

import (
	"encoding/json"
	"log"
	"time"

	"mensageria-go/internal/model"
	"mensageria-go/internal/queue"
	"mensageria-go/internal/service"
)

func StartPagamentoWorker(
	rabbit *queue.RabbitMQ,
	stripeService *service.StripeService,
) {
	log.Println("🔄 Worker de Pagamento: conectando à fila...")

	// Usar o método Consume simples
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

		// Processar pagamento
		err = stripeService.Pagar(pedido.ValorTotal)
		if err != nil {
			log.Printf("❌ Erro no pagamento do pedido %s: %v", pedido.ID, err)
			continue
		}

		log.Printf("✅ Pagamento aprovado para pedido: %s", pedido.ID)

		// Publicar evento de pagamento aprovado
		evento := map[string]interface{}{
			"evento":    "pagamento_aprovado",
			"pedido":    pedido,
			"timestamp": time.Now(),
		}
		
		body, err := json.Marshal(evento)
		if err != nil {
			log.Printf("⚠️ Erro ao serializar evento: %v", err)
			continue
		}

		// Publicar na fila de notas fiscais
		err = rabbit.Publish("notas_fiscais", body)
		if err != nil {
			log.Printf("⚠️ Erro ao publicar evento para nota fiscal: %v", err)
		} else {
			log.Printf("📤 Evento publicado para notas_fiscais: %s", pedido.ID)
		}
	}
}