package service

import (
	"errors"
	"service-acl/entity"
	"service-acl/exception/validation"
	"service-acl/model"
	"service-acl/repository"
	"service-acl/util"

	"golang.org/x/crypto/bcrypt"
)

type authServiceImpl struct {
	repository repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) AuthService {
	return &authServiceImpl{*userRepository}
}

func (service *authServiceImpl) Register(input model.RegisterRequest) (model.RegisterResponse, error) {
	//validate input
	validation.RegisterValidate(input)

	//checkunique username
	_, err := service.repository.FindByUsername(input.Username)
	if err == nil {
		return model.RegisterResponse{}, errors.New("unique")
	}

	user := entity.User{}
	user.Name = input.Name
	user.Username = input.Username
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return model.RegisterResponse{}, err
	}
	user.Password = string(passwordHash)

	//create user
	service.repository.Create(user)

	if err != nil {
		return model.RegisterResponse{}, err
	}

	//response mapping
	var response = model.RegisterResponse{}
	response.Username = input.Username

	return response, nil
}

func (service *authServiceImpl) Login(input model.LoginRequest) (model.LoginResponse, error) {
	//validate input
	validation.LoginValidate(input)

	username := input.Username
	password := input.Password

	//check user
	check_user, err := service.repository.FindByUsername(username)
	if err != nil {
		return model.LoginResponse{}, errors.New("Email Not Found")
	}

	//check login
	err = bcrypt.CompareHashAndPassword([]byte(check_user.Password), []byte(password))

	if err != nil {
		return model.LoginResponse{}, errors.New("Email/Password was wrong")
	}

	//generate token
	token, err := util.GenerateNewAccessToken(check_user)

	//response mapping
	var response = model.LoginResponse{}
	response.Username = username
	response.AccessToken = token
	response.RefreshToken = token

	return response, nil
}
