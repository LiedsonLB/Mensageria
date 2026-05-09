// internal/database/db.go
package database

import (
	"fmt"
	"log"
	"time"

	"mensageria-go/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewConnection(cfg DBConfig) (*gorm.DB, error) {
	// DSN (Data Source Name) para PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)
	
	// Conectar ao banco
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco: %v", err)
	}
	
	// Configurar pool de conexões
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	
	// Auto-migrate (cria as tabelas automaticamente)
	err = db.AutoMigrate(&model.Pedido{}, &model.StatusHistory{})
	if err != nil {
		return nil, fmt.Errorf("erro ao migrar tabelas: %v", err)
	}
	
	log.Println("✅ Conectado ao PostgreSQL e tabelas criadas/atualizadas")
	return db, nil
}