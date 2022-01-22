package controller

import (
	"net/http"
	"service-acl/exception"
	"service-acl/model"
	"service-acl/model/responder"
	"service-acl/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService *service.AuthService) AuthController {
	return AuthController{*authService}
}

func (handler *AuthController) Route(app *fiber.App) {
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
}

func (handler *AuthController) Register(c *fiber.Ctx) error {
	var input model.RegisterRequest

	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := handler.AuthService.Register(input)

	if err != nil {
		//error
		if err.Error() == "unique" {
			return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
				Code:    http.StatusBadRequest,
				Message: "Something Wrong",
				Error:   exception.NewString("Email Registered, Please Login"),
				Data:    nil,
			})
		}
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Register Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(responder.ApiResponse{
		Code:    http.StatusCreated,
		Message: "Register Success",
		Error:   nil,
		Data:    responses,
	})
}

func (handler *AuthController) Login(c *fiber.Ctx) error {
	var input model.LoginRequest

	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := handler.AuthService.Login(input)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Login Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Login Success",
		Error:   nil,
		Data:    responses,
	})
}
