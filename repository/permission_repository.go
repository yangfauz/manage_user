package repository

import "service-acl/entity"

type PermissionRepository interface {
	FindAll() (permissions []entity.Permission)
	FindById(id int) (entity.Permission, error)
	Create(permission entity.Permission)
	Update(id int, permission entity.Permission)
	Delete(id int)
}
