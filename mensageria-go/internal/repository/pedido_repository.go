// internal/repository/pedido_repository.go
package repository

import (
	"fmt"
	"mensageria-go/internal/model"
	
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PedidoRepository struct {
	db *gorm.DB
}

func NewPedidoRepository(db *gorm.DB) *PedidoRepository {
	return &PedidoRepository{db: db}
}

func (r *PedidoRepository) Create(pedido *model.Pedido) error {
	// Garantir que o ID é um UUID válido
	if pedido.ID == "" {
		pedido.ID = uuid.New().String()
	}
	
	// Validar UUID antes de criar
	if _, err := uuid.Parse(pedido.ID); err != nil {
		return fmt.Errorf("ID inválido: %v", err)
	}
	
	return r.db.Create(pedido).Error
}

func (r *PedidoRepository) FindByID(id string) (*model.Pedido, error) {
	// Validar UUID antes de buscar
	if _, err := uuid.Parse(id); err != nil {
		return nil, fmt.Errorf("ID inválido: formato UUID esperado, recebido: %s", id)
	}
	
	var pedido model.Pedido
	err := r.db.Where("id = ?", id).First(&pedido).Error
	if err != nil {
		return nil, err
	}
	return &pedido, nil
}

func (r *PedidoRepository) UpdateStatus(id string, status string) error {
	// Validar UUID
	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("ID inválido: formato UUID esperado")
	}
	
	return r.db.Model(&model.Pedido{}).
		Where("id = ?", id).
		Update("status", status).Error
}

func (r *PedidoRepository) UpdateNotaFiscal(id string, notaFiscalPath string) error {
	// Validar UUID
	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("ID inválido: formato UUID esperado")
	}
	
	return r.db.Model(&model.Pedido{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"nota_fiscal": notaFiscalPath,
			"status":      "NF_EMITIDA",
		}).Error
}

func (r *PedidoRepository) AddStatusHistory(pedidoID string, status string, mensagem string) error {
	// Validar UUID
	if _, err := uuid.Parse(pedidoID); err != nil {
		return fmt.Errorf("pedido_id inválido: formato UUID esperado")
	}
	
	history := model.StatusHistory{
		PedidoID: pedidoID,
		Status:   status,
		Mensagem: mensagem,
	}
	return r.db.Create(&history).Error
}

func (r *PedidoRepository) FindHistory(pedidoID string) ([]model.StatusHistory, error) {
	if _, err := uuid.Parse(pedidoID); err != nil {
		return nil, fmt.Errorf("pedido_id inválido: formato UUID esperado")
	}
	
	var history []model.StatusHistory
	err := r.db.Where("pedido_id = ?", pedidoID).
		Order("created_at asc").
		Find(&history).Error
	return history, err
}

func (r *PedidoRepository) FindAll(limit int, offset int) ([]model.Pedido, error) {
	var pedidos []model.Pedido
	err := r.db.Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&pedidos).Error
	return pedidos, err
}