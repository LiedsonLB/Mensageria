package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"mensageria-go/internal/config"
	"mensageria-go/internal/queue"
	"mensageria-go/internal/worker"
)

func main() {
	cfg := config.Load()
	rabbit := queue.NewRabbitMQ(cfg.RabbitURL)
	
	rabbit.SetupPubSub()
	rabbit.CriarFilaNotaFiscal()

	go worker.StartNotaFiscalWorker(rabbit)

	log.Println("🚀 Worker NOTA FISCAL iniciado")
	log.Println("📌 Pressione CTRL+C para derrubar APENAS este worker")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	
	log.Println("📴 Worker NOTA FISCAL desligado")
}