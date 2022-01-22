package repository

import (
	"service-acl/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type menuRepositoryImpl struct {
	database *gorm.DB
}

func NewMenuRepository(database *gorm.DB) MenuRepository {
	return &menuRepositoryImpl{database}
}

func (repository *menuRepositoryImpl) FindAll() (menus []entity.Menu) {

	err := repository.database.Preload("SubMenu." + clause.Associations).Where("parent_id IS NULL").Find(&menus).Error

	if err != nil {
		panic(err)
	}

	return menus
}

func (repository *menuRepositoryImpl) FindById(id int) (entity.Menu, error) {
	var menu entity.Menu

	err := repository.database.Preload("Permission").Where("id = ?", id).First(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (repository *menuRepositoryImpl) Create(menu entity.Menu) {
	err := repository.database.Create(&menu).Error

	if err != nil {
		panic(err)
	}
}

func (repository *menuRepositoryImpl) Update(id int, menu entity.Menu) {
	err := repository.database.Model(&menu).Where("id = ?", id).
		Updates(map[string]interface{}{
			"position":      menu.Position,
			"parent_id":     menu.ParentID,
			"name":          menu.Name,
			"url":           menu.Url,
			"permission_id": menu.PermissionID,
			"is_active":     menu.IsActive,
			"icon":          menu.Icon,
			"description":   menu.Description,
		}).Error

	if err != nil {
		panic(err)
	}
}

func (repository *menuRepositoryImpl) Delete(id int) {
	var menu entity.Menu
	err := repository.database.Where("id = ?", id).Delete(&menu).Error

	if err != nil {
		panic(err)
	}
}
