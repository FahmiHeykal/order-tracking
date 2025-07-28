package repository

import (
	"order-tracking/internal/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) FindByID(id uint) (*model.Order, error) {
	var order model.Order
	err := r.db.Preload("User").First(&order, id).Error
	return &order, err
}

func (r *OrderRepository) FindByUserID(userID uint) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) FindAll() ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Preload("User").Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) UpdateStatus(id uint, status model.OrderStatus, changedBy uint) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&model.Order{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		tx.Rollback()
		return err
	}

	history := model.OrderStatusHistory{
		OrderID:   id,
		Status:    status,
		ChangedBy: changedBy,
	}

	if err := tx.Create(&history).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
