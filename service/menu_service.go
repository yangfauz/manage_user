package service

import "service-acl/model"

type MenuService interface {
	List() ([]model.GetAllMenuResponse, error)
	Detail(id int) (model.GetMenuResponse, error)
	Create(input model.CreateMenuRequest) (model.CreateMenuResponse, error)
	Edit(id int, input model.UpdateMenuRequest) (model.UpdateMenuResponse, error)
	Delete(id int) error
}
