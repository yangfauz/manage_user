package service

import (
	"service-acl/entity"
	"service-acl/exception/validation"
	"service-acl/model"
	"service-acl/repository"
)

type permissionServiceImpl struct {
	repository repository.PermissionRepository
}

func NewPermissionService(permissionRepository *repository.PermissionRepository) PermissionService {
	return &permissionServiceImpl{*permissionRepository}
}

func (service *permissionServiceImpl) List() ([]model.GetPermissionResponse, error) {
	//get data permission
	permissions := service.repository.FindAll()

	// mapping response
	responses := []model.GetPermissionResponse{}

	for _, permission := range permissions {
		response := model.GetPermissionResponse{}
		response.ID = permission.ID
		response.Name = permission.Name
		response.Description = permission.Description

		responses = append(responses, response)
	}

	return responses, nil
}

func (service *permissionServiceImpl) Detail(id int) (model.GetPermissionResponse, error) {
	permission, err := service.repository.FindById(id)

	if err != nil {
		return model.GetPermissionResponse{}, err
	}

	//mapping response
	var response = model.GetPermissionResponse{}
	response.ID = permission.ID
	response.Name = permission.Name
	response.Description = permission.Description

	return response, nil
}

func (service *permissionServiceImpl) Create(input model.CreatePermissionRequest) (model.CreatePermissionResponse, error) {
	//validate input
	validation.CreatePermissionValidate(input)

	permission := entity.Permission{}
	permission.Name = input.Name
	permission.Description = input.Description

	//update permission
	service.repository.Create(permission)

	//mapping response
	var response = model.CreatePermissionResponse{}
	response.Name = permission.Name

	return response, nil
}

func (service *permissionServiceImpl) Edit(id int, input model.UpdatePermissionRequest) (model.UpdatePermissionResponse, error) {
	//validate input
	validation.UpdatePermissionValidate(input)

	//check permission
	_, err := service.repository.FindById(id)

	if err != nil {
		return model.UpdatePermissionResponse{}, err
	}

	permission := entity.Permission{}
	permission.Name = input.Name
	permission.Description = input.Description

	//update permission
	service.repository.Update(id, permission)

	//mapping response
	var response = model.UpdatePermissionResponse{}
	response.Name = permission.Name

	return response, nil
}

func (service *permissionServiceImpl) Delete(id int) error {
	_, err := service.repository.FindById(id)

	if err != nil {
		return err
	}

	//delete permission
	service.repository.Delete(id)

	return nil
}
