package service

import (
	"service-acl/model"
)

type AuthService interface {
	Register(input model.RegisterRequest) (model.RegisterResponse, error)
	Login(input model.LoginRequest) (model.LoginResponse, error)
}
