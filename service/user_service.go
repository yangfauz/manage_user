package service

import "service-acl/model"

type UserService interface {
	List() ([]model.GetUserResponse, error)
	Detail(id int) (model.GetUserResponse, error)
	Edit(id int, input model.UpdateUserRequest) (model.UpdateUserResponse, error)
}
