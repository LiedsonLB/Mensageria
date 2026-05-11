// cmd/worker-email/main.go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"mensageria-go/internal/config"
	"mensageria-go/internal/database"
	"mensageria-go/internal/queue"
	"mensageria-go/internal/service"
	"mensageria-go/internal/worker"
)

func main() {
	cfg := config.Load()
	
	dbConfig := database.DBConfig{
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
	
	rabbit := queue.NewRabbitMQ(cfg.RabbitURL)
	
	rabbit.SetupPubSub()
	rabbit.CriarFilaDanfeGerado()

	emailService := service.NewEmailService(cfg)

	go worker.StartEmailWorker(rabbit, emailService, db)

	log.Println("🚀 Worker EMAIL iniciado")
	log.Println("📌 Pressione CTRL+C para derrubar APENAS este worker")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	
	log.Println("📴 Worker EMAIL desligado")
}