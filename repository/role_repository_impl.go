package repository

import (
	"service-acl/entity"

	"gorm.io/gorm"
)

type roleRepositoryImpl struct {
	database *gorm.DB
}

func NewRoleRepository(database *gorm.DB) RoleRepository {
	return &roleRepositoryImpl{database}
}

func (repository *roleRepositoryImpl) FindAll() (roles []entity.Role) {

	err := repository.database.Preload("Permission").Find(&roles).Error

	if err != nil {
		panic(err)
	}

	return roles
}

func (repository *roleRepositoryImpl) FindById(id int) (entity.Role, error) {
	var role entity.Role
	err := repository.database.Preload("Permission").Where("id = ?", id).First(&role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}

func (repository *roleRepositoryImpl) Create(role entity.Role) {
	err := repository.database.Create(&role).Error

	if err != nil {
		panic(err)
	}
}

func (repository *roleRepositoryImpl) Update(id int, role entity.Role) {
	err := repository.database.Model(&role).Where("id = ?", id).
		Updates(map[string]interface{}{"name": role.Name, "description": role.Description}).Error

	if err != nil {
		panic(err)
	}
}

func (repository *roleRepositoryImpl) Delete(id int) {
	var role entity.Role
	err := repository.database.Where("id = ?", id).Delete(&role).Error

	if err != nil {
		panic(err)
	}
}

func (repository *roleRepositoryImpl) RolePermission(role entity.Role, permission_id []int) {

	permissions := make([]*entity.Permission, len(permission_id))

	for i, id := range permission_id {
		permissions[i] = &entity.Permission{ID: uint(id)}
	}

	err := repository.database.Model(&role).Association("Permission").Replace(permissions)

	if err != nil {
		panic(err)
	}
}
