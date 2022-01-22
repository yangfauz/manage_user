package service

import "service-acl/model"

type ProfileService interface {
	ProfileDetail(user_id int) (model.GetUserResponse, error)
	ProfileRole(user_id int) ([]model.GetRoleResponse, error)
	ProfilePermission(user_id int) ([]model.GetPermissionResponse, error)
}
