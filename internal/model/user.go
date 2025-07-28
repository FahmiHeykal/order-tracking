package model

import "gorm.io/gorm"

type Role string

const (
	RoleUser   Role = "user"
	RoleAdmin  Role = "admin"
	RoleDriver Role = "driver"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     Role   `gorm:"type:varchar(10);not null;default:'user'"`
}
