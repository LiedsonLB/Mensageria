// internal/handler/pedido_handler.go
package handler

import (
	"encoding/json"
	"log"

	"mensageria-go/internal/model"
	"mensageria-go/internal/queue"
	"mensageria-go/internal/repository"

	"github.com/gin-gonic/gin"
)

type PedidoHandler struct {
	rabbit   *queue.RabbitMQ
	repo     *repository.PedidoRepository
}

func NewPedidoHandler(rabbit *queue.RabbitMQ, repo *repository.PedidoRepository) *PedidoHandler {
	return &PedidoHandler{
		rabbit: rabbit,
		repo:   repo,
	}
}

func (h *PedidoHandler) CriarPedido(c *gin.Context) {
	var pedido model.Pedido
	
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	pedido.ValorTotal = 0
	for i := range pedido.ProdutosList {
		pedido.ProdutosList[i].Subtotal = pedido.ProdutosList[i].Preco * float64(pedido.ProdutosList[i].Quantidade)
		pedido.ValorTotal += pedido.ProdutosList[i].Subtotal
	}
	
	pedido.Status = "PENDENTE"
	
	if err := h.repo.Create(&pedido); err != nil {
		log.Printf("❌ Erro ao salvar pedido no banco: %v", err)
		c.JSON(500, gin.H{"error": "Erro ao salvar pedido"})
		return
	}
	
	h.repo.AddStatusHistory(pedido.ID, "PENDENTE", "Pedido criado e aguardando processamento")
	
	body, err := json.Marshal(pedido)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao serializar pedido"})
		return
	}
	
	if err := h.rabbit.Publish("pagamentos", body); err != nil {
		log.Printf("⚠️ Erro ao publicar na fila: %v", err)
	}
	
	c.JSON(201, pedido)
}

func (h *PedidoHandler) BuscarPedido(c *gin.Context) {
	id := c.Param("id")
	
	pedido, err := h.repo.FindByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Pedido não encontrado"})
		return
	}
	
	c.JSON(200, pedido)
}

func (h *PedidoHandler) BuscarPedidos(c *gin.Context) {
    pedidos, err := h.repo.FindAll(100, 0)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    result := []gin.H{}
    for _, p := range pedidos {
        produtos, _ := p.GetProdutosList()
        result = append(result, gin.H{
            "id":          p.ID,
            "cliente":     p.Cliente,
            "documento":   p.Documento,
            "email":       p.Email,
            "valor_total": p.ValorTotal,
            "status":      p.Status,
            "produtos":    produtos,
            "created_at":  p.CreatedAt,
            "updated_at":  p.UpdatedAt,
        })
    }
    
    c.JSON(200, result)
}

func (h *PedidoHandler) BuscarStatus(c *gin.Context) {
	id := c.Param("id")
	
	pedido, err := h.repo.FindByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Pedido não encontrado"})
		return
	}
	
	history, _ := h.repo.FindHistory(id)
	
	c.JSON(200, gin.H{
		"id":      pedido.ID,
		"status":  pedido.Status,
		"history": history,
	})
}