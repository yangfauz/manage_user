package repository

import "service-acl/entity"

type RoleRepository interface {
	FindAll() (roles []entity.Role)
	FindById(id int) (entity.Role, error)
	Create(role entity.Role)
	Update(id int, role entity.Role)
	Delete(id int)
	RolePermission(role entity.Role, permission_id []int)
}
