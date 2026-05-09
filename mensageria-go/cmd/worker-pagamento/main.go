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
	rabbit.CriarFilaPagamento()

	stripeService := service.NewStripeService(cfg)

	go worker.StartPagamentoWorker(rabbit, stripeService, db)

	log.Println("🚀 Worker PAGAMENTO iniciado")
	log.Println("📌 Pressione CTRL+C para derrubar APENAS este worker")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("📴 Worker PAGAMENTO desligado")
}