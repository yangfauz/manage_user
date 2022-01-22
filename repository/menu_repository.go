package repository

import "service-acl/entity"

type MenuRepository interface {
	FindAll() (menus []entity.Menu)
	FindById(id int) (entity.Menu, error)
	Create(menu entity.Menu)
	Update(id int, menu entity.Menu)
	Delete(id int)
}
