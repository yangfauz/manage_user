package service

import "service-acl/model"

type AssignService interface {
	UserRole(input model.UserRoleRequest) (model.UserRoleResponse, error)
	RolePermission(input model.RolePermissionRequest) (model.RolePermissionResponse, error)
}
