package handler

import (
	"net/http"

	"mensageria-go/internal/queue"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	rabbit *queue.RabbitMQ
}

func NewDashboardHandler(rabbit *queue.RabbitMQ) *DashboardHandler {
	return &DashboardHandler{rabbit: rabbit}
}

func (h *DashboardHandler) GetQueueStatus(c *gin.Context) {
	status := map[string]interface{}{
		"pagamentos":     getQueueInfo(h.rabbit, "pagamentos"),
		"pagamentos_dlq": getQueueInfo(h.rabbit, "pagamentos_dlq"),
		"emails":         getQueueInfo(h.rabbit, "emails"),
		"workers": map[string]interface{}{
			"pagamento": "running",
			"email":     "running",
		},
	}
	
	c.JSON(http.StatusOK, status)
}

func getQueueInfo(rabbit *queue.RabbitMQ, queueName string) map[string]interface{} {
	// Para implementar métricas reais, você precisaria acessar o canal do RabbitMQ
	// Mas como o queue.RabbitMQ não expõe o channel, vamos retornar info básica
	
	// Se quiser implementar completamente, adicione um método no queue.RabbitMQ:
	// rabbit.GetQueueStats(queueName)
	
	return map[string]interface{}{
		"name":    queueName,
		"status":  "active",
		"message": "Implemente o método GetQueueStats no queue.RabbitMQ para métricas reais",
	}
}