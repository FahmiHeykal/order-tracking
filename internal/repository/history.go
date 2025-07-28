package repository

import (
	"order-tracking/internal/model"

	"gorm.io/gorm"
)

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{db: db}
}

func (r *HistoryRepository) FindByOrderID(orderID uint) ([]model.OrderStatusHistory, error) {
	var histories []model.OrderStatusHistory
	err := r.db.Where("order_id = ?", orderID).Preload("User").Order("created_at desc").Find(&histories).Error
	return histories, err
}
