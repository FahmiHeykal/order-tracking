package model

import "gorm.io/gorm"

type OrderStatusHistory struct {
	gorm.Model
	OrderID   uint        `gorm:"not null"`
	Order     Order       `gorm:"foreignKey:OrderID"`
	Status    OrderStatus `gorm:"type:varchar(20);not null"`
	ChangedBy uint        `gorm:"not null"`
	User      User        `gorm:"foreignKey:ChangedBy"`
}
