package main

import (
	"fmt"
	"mensageria-go/internal/config"
	"mensageria-go/internal/handler"
	"mensageria-go/internal/queue"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	rabbit := queue.NewRabbitMQ(cfg.RabbitURL)

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		
		allowedOrigins := map[string]bool{
			"http://localhost:3000": true,
			"http://localhost:3001": true,
			"http://localhost:5173": true,
		}
		
		if allowedOrigins[origin] {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001")
		}
		
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	h := handler.NewPedidoHandler(rabbit)

	router.POST("/pedido", h.CriarPedido)

	router.Static("/notas-fiscais", "./notas-fiscais")
	
	router.GET("/download/:id", func(c *gin.Context) {
		id := c.Param("id")
		filePath := fmt.Sprintf("notas-fiscais/%s.xml", id)
		c.File(filePath)
	})

	router.Run(":8080")
}