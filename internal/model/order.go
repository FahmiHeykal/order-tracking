package model

import "gorm.io/gorm"

type OrderStatus string

const (
	StatusPending   OrderStatus = "Pending"
	StatusProcessed OrderStatus = "Diproses"
	StatusShipped   OrderStatus = "Dikirim"
	StatusCompleted OrderStatus = "Selesai"
)

func (s OrderStatus) IsValid() bool {
	switch s {
	case StatusPending, StatusProcessed, StatusShipped, StatusCompleted:
		return true
	default:
		return false
	}
}

type Order struct {
	gorm.Model
	UserID      uint        `gorm:"not null"`
	User        User        `gorm:"foreignKey:UserID"`
	Status      OrderStatus `gorm:"type:varchar(20);not null;default:'Pending'"`
	Description string
}
