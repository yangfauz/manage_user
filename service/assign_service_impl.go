package service

import (
	"service-acl/model"
	"service-acl/repository"
)

type assignServiceImpl struct {
	repositoryUser repository.UserRepository
	repositoryRole repository.RoleRepository
}

func NewAssignService(userRepository *repository.UserRepository, roleRepository *repository.RoleRepository) AssignService {
	return &assignServiceImpl{
		*userRepository,
		*roleRepository,
	}
}

func (service *assignServiceImpl) UserRole(input model.UserRoleRequest) (model.UserRoleResponse, error) {
	user, err := service.repositoryUser.FindById(input.UserID)

	if err != nil {
		return model.UserRoleResponse{}, err
	}

	//update user role
	service.repositoryUser.UserRole(user, input.RoleID)

	//mapping response
	var response = model.UserRoleResponse{}
	response.UserID = input.UserID

	return response, nil
}

func (service *assignServiceImpl) RolePermission(input model.RolePermissionRequest) (model.RolePermissionResponse, error) {
	role, err := service.repositoryRole.FindById(input.RoleID)

	if err != nil {
		return model.RolePermissionResponse{}, err
	}

	//update user role
	service.repositoryRole.RolePermission(role, input.PermissionID)

	//mapping response
	var response = model.RolePermissionResponse{}
	response.RoleID = input.RoleID

	return response, nil
}
