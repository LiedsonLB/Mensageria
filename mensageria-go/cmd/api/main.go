package main

import (
	"fmt"
	"log"

	"mensageria-go/internal/config"
	"mensageria-go/internal/database"
	"mensageria-go/internal/handler"
	"mensageria-go/internal/queue"
	"mensageria-go/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	dbConfig := struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	}
	
	db, err := database.NewConnection(dbConfig)
	if err != nil {
		log.Fatal("❌ Erro ao conectar ao banco:", err)
	}
	
	pedidoRepo := repository.NewPedidoRepository(db)
	rabbit := queue.NewRabbitMQ(cfg.RabbitURL)
	pedidoHandler := handler.NewPedidoHandler(rabbit, pedidoRepo)

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

	router.POST("/pedido", pedidoHandler.CriarPedido)
	router.GET("/pedido/:id", pedidoHandler.BuscarPedido)
	router.GET("/pedidos", pedidoHandler.BuscarPedidos)
	router.GET("/pedido/:id/status", pedidoHandler.BuscarStatus)
	router.Static("/notas-fiscais", "./notas-fiscais")
	router.Static("/notas-fiscais-pdf", "./notas-fiscais-pdf")
	router.GET("/download/:id", func(c *gin.Context) {
		id := c.Param("id")
		filePath := fmt.Sprintf("notas-fiscais/%s.xml", id)
		c.File(filePath)
	})

	log.Println("🚀 API iniciada na porta 8080")
	router.Run(":8080")
}