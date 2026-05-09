package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RabbitURL string
	StripeKey string

	// EMAIL SMTP
	EmailUser string
	EmailPass string
	EmailHost string
	EmailPort string
}

// ======================================
// LOAD CONFIG
// ======================================
func Load() Config {

	// carrega .env automaticamente
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env não encontrado (usando variáveis do sistema)")
	}

	cfg := Config{
		// RabbitMQ
		RabbitURL: getEnv(
			"RABBIT_URL",
			"amqp://admin:admin123@localhost:5672/",
		),

		// Stripe (OBRIGATÓRIO)
		StripeKey: mustGetEnv("STRIPE_KEY"),

		// Email SMTP
		EmailUser: mustGetEnv("EMAIL_USER"),
		EmailPass: mustGetEnv("EMAIL_PASS"),
		EmailHost: getEnv("EMAIL_HOST", "smtp.gmail.com"),
		EmailPort: getEnv("EMAIL_PORT", "587"),
	}

	log.Println("✅ Configurações carregadas")

	return cfg
}

// ======================================
// VAR OPCIONAL
// ======================================
func getEnv(key, fallback string) string {

	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	return val
}

// ======================================
// VAR OBRIGATÓRIA
// ======================================
func mustGetEnv(key string) string {

	val := os.Getenv(key)

	if val == "" {
		log.Fatalf("❌ Variável obrigatória não definida: %s", key)
	}

	return val
}