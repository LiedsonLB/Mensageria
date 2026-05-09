// internal/handler/pedido_handler.go
package handler

import (
	"encoding/json"
	
	"mensageria-go/internal/model"
	"mensageria-go/internal/queue"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PedidoHandler struct {
	rabbit *queue.RabbitMQ
}

func NewPedidoHandler(rabbit *queue.RabbitMQ) *PedidoHandler {
	return &PedidoHandler{rabbit: rabbit}
}

func (h *PedidoHandler) CriarPedido(c *gin.Context) {
	var pedido model.Pedido
	
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	// Gerar ID automaticamente
	if pedido.ID == "" {
		pedido.ID = uuid.New().String()
	}
	
	// Calcular valor total e subtotal de cada produto
	pedido.ValorTotal = 0
	for i := range pedido.Produtos {
		// Calcular subtotal do produto
		pedido.Produtos[i].Subtotal = pedido.Produtos[i].Preco * float64(pedido.Produtos[i].Quantidade)
		// Acumular no valor total
		pedido.ValorTotal += pedido.Produtos[i].Subtotal
	}
	
	pedido.Status = "PENDENTE"
	
	// Converter para JSON
	body, err := json.Marshal(pedido)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao serializar pedido"})
		return
	}
	
	// Publicar na fila
	h.rabbit.Publish("pagamentos", body)
	
	c.JSON(201, pedido)
}