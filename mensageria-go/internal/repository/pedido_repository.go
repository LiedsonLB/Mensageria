// internal/repository/pedido_repository.go
package repository

import (
	"mensageria-go/internal/model"
	
	"gorm.io/gorm"
)

type PedidoRepository struct {
	db *gorm.DB
}

func NewPedidoRepository(db *gorm.DB) *PedidoRepository {
	return &PedidoRepository{db: db}
}

func (r *PedidoRepository) Create(pedido *model.Pedido) error {
	return r.db.Create(pedido).Error
}

func (r *PedidoRepository) FindByID(id string) (*model.Pedido, error) {
	var pedido model.Pedido
	err := r.db.First(&pedido, "id = ?", id).Error
	return &pedido, err
}

func (r *PedidoRepository) UpdateStatus(id string, status string) error {
	return r.db.Model(&model.Pedido{}).
		Where("id = ?", id).
		Update("status", status).Error
}

func (r *PedidoRepository) UpdateNotaFiscal(id string, notaFiscalPath string) error {
	return r.db.Model(&model.Pedido{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"nota_fiscal": notaFiscalPath,
			"status":      "NF_EMITIDA",
		}).Error
}

func (r *PedidoRepository) AddStatusHistory(pedidoID string, status string, mensagem string) error {
	history := model.StatusHistory{
		PedidoID: pedidoID,
		Status:   status,
		Mensagem: mensagem,
	}
	return r.db.Create(&history).Error
}

func (r *PedidoRepository) FindHistory(pedidoID string) ([]model.StatusHistory, error) {
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