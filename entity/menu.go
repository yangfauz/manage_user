package entity

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	ID           uint       `gorm:"primaryKey"`
	Position     int        `gorm:"not null"`
	ParentID     *uint      `gorm:"default:null"`
	Name         string     `gorm:"size:100;not null"`
	Url          string     `gorm:"not null;type:text"`
	PermissionID *uint      `gorm:"default:null"`
	IsActive     bool       `gorm:"not null"`
	Icon         string     `gorm:"size:100;not null"`
	Description  *string    `gorm:"default:null;type:text"`
	CreatedBy    *uint      `gorm:"default:null"`
	UpdatedBy    *uint      `gorm:"default:null"`
	DeletedBy    *uint      `gorm:"default:null"`
	CreatedAt    *time.Time `gorm:"default:null"`
	UpdatedAt    *time.Time `gorm:"default:null"`
	DeletedAt    *time.Time `gorm:"default:null"`
	SubMenu      []Menu     `gorm:"foreignKey:ParentID"`
	Permission   Permission `gorm:"foreignKey:PermissionID"`
}
