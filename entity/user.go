package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"size:200;not null"`
	Username  string     `gorm:"size:100;not null;index:idx_name,unique"`
	Password  string     `gorm:"size:250;not null"`
	CreatedBy *uint      `gorm:"default:null"`
	UpdatedBy *uint      `gorm:"default:null"`
	DeletedBy *uint      `gorm:"default:null"`
	CreatedAt *time.Time `gorm:"default:null"`
	UpdatedAt *time.Time `gorm:"default:null"`
	DeletedAt *time.Time `gorm:"default:null"`
	Role      []*Role    `gorm:"many2many:has_user_roles;"`
}
