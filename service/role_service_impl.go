package service

import (
	"service-acl/entity"
	"service-acl/exception/validation"
	"service-acl/model"
	"service-acl/repository"
)

type roleServiceImpl struct {
	repository repository.RoleRepository
}

func NewRoleService(roleRepository *repository.RoleRepository) RoleService {
	return &roleServiceImpl{*roleRepository}
}

func (service *roleServiceImpl) List() ([]model.GetRoleResponse, error) {
	//get data role
	roles := service.repository.FindAll()

	// mapping response
	responses := []model.GetRoleResponse{}

	for _, role := range roles {
		response := model.GetRoleResponse{}
		response.ID = role.ID
		response.Name = role.Name
		response.Description = role.Description

		responses = append(responses, response)
	}

	return responses, nil
}

func (service *roleServiceImpl) Detail(id int) (model.GetRoleResponse, error) {
	role, err := service.repository.FindById(id)

	if err != nil {
		return model.GetRoleResponse{}, err
	}

	//mapping response
	responses_permission := []model.GetPermissionResponse{}
	for _, permission := range role.Permission {
		response_permission := model.GetPermissionResponse{}
		response_permission.ID = permission.ID
		response_permission.Name = permission.Name

		responses_permission = append(responses_permission, response_permission)
	}

	var response = model.GetRoleResponse{}
	response.ID = role.ID
	response.Name = role.Name
	response.Description = role.Description
	response.Permissions = responses_permission

	return response, nil
}

func (service *roleServiceImpl) Create(input model.CreateRoleRequest) (model.CreateRoleResponse, error) {
	//validate input
	validation.CreateRoleValidate(input)

	role := entity.Role{}
	role.Name = input.Name
	role.Description = input.Description

	//update role
	service.repository.Create(role)

	//mapping response
	var response = model.CreateRoleResponse{}
	response.Name = role.Name

	return response, nil
}

func (service *roleServiceImpl) Edit(id int, input model.UpdateRoleRequest) (model.UpdateRoleResponse, error) {
	//validate input
	validation.UpdateRoleValidate(input)

	//check role
	_, err := service.repository.FindById(id)

	if err != nil {
		return model.UpdateRoleResponse{}, err
	}

	role := entity.Role{}
	role.Name = input.Name
	role.Description = input.Description

	//update role
	service.repository.Update(id, role)

	//mapping response
	var response = model.UpdateRoleResponse{}
	response.Name = role.Name

	return response, nil
}

func (service *roleServiceImpl) Delete(id int) error {
	_, err := service.repository.FindById(id)

	if err != nil {
		return err
	}

	//delete role
	service.repository.Delete(id)

	return nil
}
