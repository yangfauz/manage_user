package service

import (
	"service-acl/model"
	"service-acl/repository"
)

type profileServiceImpl struct {
	repositoryUser repository.UserRepository
}

func NewProfileService(userRepository *repository.UserRepository) ProfileService {
	return &profileServiceImpl{*userRepository}
}

func (service *profileServiceImpl) ProfileDetail(user_id int) (model.GetUserResponse, error) {
	user, err := service.repositoryUser.FindById(user_id)

	if err != nil {
		return model.GetUserResponse{}, err
	}

	//mapping response
	responses_role := []model.GetRoleResponse{}
	for _, role := range user.Role {
		response_role := model.GetRoleResponse{}
		response_role.ID = role.ID
		response_role.Name = role.Name

		responses_role = append(responses_role, response_role)
	}

	var response = model.GetUserResponse{}
	response.ID = user.ID
	response.Name = user.Name
	response.Username = user.Username
	response.Roles = responses_role

	return response, nil
}

func (service *profileServiceImpl) ProfileRole(user_id int) ([]model.GetRoleResponse, error) {
	user, err := service.repositoryUser.FindById(user_id)

	if err != nil {
		return []model.GetRoleResponse{}, err
	}

	//mapping response
	responses_role := []model.GetRoleResponse{}
	for _, role := range user.Role {
		response_role := model.GetRoleResponse{}
		response_role.ID = role.ID
		response_role.Name = role.Name

		responses_role = append(responses_role, response_role)
	}

	return responses_role, nil
}

func (service *profileServiceImpl) ProfilePermission(user_id int) ([]model.GetPermissionResponse, error) {
	user, err := service.repositoryUser.FindById(user_id)

	if err != nil {
		return []model.GetPermissionResponse{}, err
	}

	//mapping response
	responses_role := []model.GetRoleResponse{}
	responses_permission := []model.GetPermissionResponse{}
	for _, role := range user.Role {
		response_role := model.GetRoleResponse{}
		response_role.ID = role.ID
		response_role.Name = role.Name

		responses_role = append(responses_role, response_role)

		for _, permission := range role.Permission {
			response_permission := model.GetPermissionResponse{}
			response_permission.ID = permission.ID
			response_permission.Name = permission.Name

			responses_permission = append(responses_permission, response_permission)
		}
	}

	return responses_permission, nil
}
