package service

import "service-acl/model"

type PermissionService interface {
	List() ([]model.GetPermissionResponse, error)
	Detail(id int) (model.GetPermissionResponse, error)
	Create(input model.CreatePermissionRequest) (model.CreatePermissionResponse, error)
	Edit(id int, input model.UpdatePermissionRequest) (model.UpdatePermissionResponse, error)
	Delete(id int) error
}
