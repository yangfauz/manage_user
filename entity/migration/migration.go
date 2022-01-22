package migration

import (
	"service-acl/entity"

	"gorm.io/gorm"
)

func Migration(database *gorm.DB) {
	database.AutoMigrate(
		&entity.Role{},
		&entity.Permission{},
		&entity.Menu{},
		&entity.User{},
	)
}
