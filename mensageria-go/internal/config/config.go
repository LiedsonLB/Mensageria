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

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// ======================================
// LOAD CONFIG
// ======================================
func Load() Config {

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env não encontrado (usando variáveis do sistema)")
	}

	cfg := Config{
		RabbitURL: getEnv(
			"RABBIT_URL",
			"amqp://admin:admin123@localhost:5672/",
		),

		StripeKey: mustGetEnv("STRIPE_KEY"),

		EmailUser: mustGetEnv("EMAIL_USER"),
		EmailPass: mustGetEnv("EMAIL_PASS"),
		EmailHost: getEnv("EMAIL_HOST", "smtp.gmail.com"),
		EmailPort: getEnv("EMAIL_PORT", "587"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "admin"),
		DBPassword: getEnv("DB_PASSWORD", "admin123"),
		DBName:     getEnv("DB_NAME", "mensageria"),
	}

	log.Println("✅ Configurações carregadas")
	log.Printf("📦 DB: %s:%s/%s", cfg.DBHost, cfg.DBPort, cfg.DBName)
	log.Printf("🐇 RabbitMQ: %s", cfg.RabbitURL)

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