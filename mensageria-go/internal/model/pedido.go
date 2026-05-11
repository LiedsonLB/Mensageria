package model

import (
	"log"
	"encoding/json"
	"github.com/google/uuid"
	"time"
	"gorm.io/gorm"
)

type Pedido struct {
	ID         string         `gorm:"primaryKey;type:uuid" json:"id"`
	Cliente    string         `gorm:"not null;size:255" json:"cliente"`
	Documento  string         `gorm:"not null;size:20" json:"documento"`
	Email      string         `gorm:"not null;size:255;index" json:"email"`
	ValorTotal float64        `gorm:"type:decimal(10,2)" json:"valor_total"`
	Status     string         `gorm:"default:'PENDENTE';size:50;index" json:"status"`
	Produtos   string         `gorm:"type:jsonb" json:"-"`
	NotaFiscal string         `gorm:"size:500" json:"nota_fiscal,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	
	ProdutosList []ProdutoItem `gorm:"-" json:"produtos"`
}

// BeforeCreate - Hook executado antes de criar no banco
func (p *Pedido) BeforeCreate(tx *gorm.DB) error {
	// Gerar UUID se não existir
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	
	// Status padrão
	if p.Status == "" {
		p.Status = "PENDENTE"
	}
	
	// Converter ProdutosList para JSON string
	if len(p.ProdutosList) > 0 {
		produtosJSON, err := json.Marshal(p.ProdutosList)
		if err != nil {
			return err
		}
		p.Produtos = string(produtosJSON)
	}
	
	return nil
}

// ProdutosList retorna a lista de produtos (não é campo do banco)
func (p *Pedido) GetProdutosList() ([]ProdutoItem, error) {
	var produtos []ProdutoItem
	if p.Produtos != "" {
		err := json.Unmarshal([]byte(p.Produtos), &produtos)
		if err != nil {
			log.Printf("❌ Erro ao deserializar produtos: %v", err)
			return produtos, err
		}
		log.Printf("✅ Produtos carregados: %d itens", len(produtos))
	}
	return produtos, nil
}

// AfterFind - Hook executado após buscar do banco
func (p *Pedido) AfterFind(tx *gorm.DB) error {
	// Converter JSON string para ProdutosList
	if p.Produtos != "" {
		err := json.Unmarshal([]byte(p.Produtos), &p.ProdutosList)
		if err != nil {
			return err
		}
	}
	return nil
}

// BeforeUpdate - Hook executado antes de atualizar
func (p *Pedido) BeforeUpdate(tx *gorm.DB) error {
	// Converter ProdutosList para JSON string se houver mudanças
	if len(p.ProdutosList) > 0 {
		produtosJSON, err := json.Marshal(p.ProdutosList)
		if err != nil {
			return err
		}
		p.Produtos = string(produtosJSON)
	}
	return nil
}

// TableName especifica o nome da tabela no PostgreSQL
func (Pedido) TableName() string {
	return "pedidos"
}

// StatusHistory - Histórico de status do pedido
type StatusHistory struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PedidoID  string    `gorm:"type:uuid;index" json:"pedido_id"`
	Status    string    `gorm:"size:50;not null" json:"status"`
	Mensagem  string    `gorm:"type:text" json:"mensagem"`
	CreatedAt time.Time `json:"created_at"`
}

func (StatusHistory) TableName() string {
	return "status_history"
}