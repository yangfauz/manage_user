package entity

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID          uint       `gorm:"primaryKey"`
	Name        string     `gorm:"size:100;not null;index:idx_name,unique"`
	Description *string    `gorm:"default:null;type:text"`
	CreatedBy   *uint      `gorm:"default:null"`
	UpdatedBy   *uint      `gorm:"default:null"`
	DeletedBy   *uint      `gorm:"default:null"`
	CreatedAt   *time.Time `gorm:"default:null"`
	UpdatedAt   *time.Time `gorm:"default:null"`
	DeletedAt   *time.Time `gorm:"default:null"`
	Role        []*Role    `gorm:"many2many:has_role_permissions;"`
}
