package service

import "service-acl/model"

type RoleService interface {
	List() ([]model.GetRoleResponse, error)
	Detail(id int) (model.GetRoleResponse, error)
	Create(input model.CreateRoleRequest) (model.CreateRoleResponse, error)
	Edit(id int, input model.UpdateRoleRequest) (model.UpdateRoleResponse, error)
	Delete(id int) error
}
