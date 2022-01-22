package repository

import (
	"service-acl/entity"

	"gorm.io/gorm"
)

type permissionRepositoryImpl struct {
	database *gorm.DB
}

func NewPermissionRepository(database *gorm.DB) PermissionRepository {
	return &permissionRepositoryImpl{database}
}

func (repository *permissionRepositoryImpl) FindAll() (permissions []entity.Permission) {

	err := repository.database.Find(&permissions).Error

	if err != nil {
		panic(err)
	}

	return permissions
}

func (repository *permissionRepositoryImpl) FindById(id int) (entity.Permission, error) {
	var permission entity.Permission

	err := repository.database.Where("id = ?", id).First(&permission).Error

	if err != nil {
		return permission, err
	}

	return permission, nil
}

func (repository *permissionRepositoryImpl) Create(permission entity.Permission) {
	err := repository.database.Create(&permission).Error

	if err != nil {
		panic(err)
	}
}

func (repository *permissionRepositoryImpl) Update(id int, permission entity.Permission) {
	err := repository.database.Model(&permission).Where("id = ?", id).
		Updates(map[string]interface{}{"name": permission.Name, "description": permission.Description}).Error

	if err != nil {
		panic(err)
	}
}

func (repository *permissionRepositoryImpl) Delete(id int) {
	var permission entity.Permission
	err := repository.database.Where("id = ?", id).Delete(&permission).Error

	if err != nil {
		panic(err)
	}
}
