package service

import (
	"service-acl/entity"
	"service-acl/exception/validation"
	"service-acl/model"
	"service-acl/repository"
)

type userServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{*userRepository}
}

func (service *userServiceImpl) List() ([]model.GetUserResponse, error) {
	//get data user
	users := service.repository.FindAll()

	// mapping response
	responses := []model.GetUserResponse{}

	for _, user := range users {
		response := model.GetUserResponse{}
		response.ID = user.ID
		response.Name = user.Name
		response.Username = user.Username

		responses = append(responses, response)
	}

	return responses, nil
}

func (service *userServiceImpl) Detail(id int) (model.GetUserResponse, error) {
	user, err := service.repository.FindById(id)

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

func (service *userServiceImpl) Edit(id int, input model.UpdateUserRequest) (model.UpdateUserResponse, error) {
	//validate input
	validation.UpdateUserValidate(input)

	//check user
	_, err := service.repository.FindById(id)

	if err != nil {
		return model.UpdateUserResponse{}, err
	}

	user := entity.User{}
	user.Name = input.Name

	//update user
	service.repository.Update(id, user)
	if err != nil {
		return model.UpdateUserResponse{}, err
	}

	//response mapping
	var response = model.UpdateUserResponse{}
	response.Name = input.Name

	return response, nil
}
