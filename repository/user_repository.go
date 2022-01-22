package repository

import "service-acl/entity"

type UserRepository interface {
	FindAll() (users []entity.User)
	FindByUsername(username string) (entity.User, error)
	FindById(id int) (entity.User, error)
	Create(user entity.User)
	Update(id int, user entity.User)
	UserRole(user entity.User, role_id []int)
}
