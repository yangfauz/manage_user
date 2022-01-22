package controller

import (
	"net/http"
	"service-acl/exception"
	"service-acl/middleware"
	"service-acl/model"
	"service-acl/model/responder"
	"service-acl/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{*userService}
}

func (handler *UserController) Route(app *fiber.App) {
	app.Get("/user", middleware.JWTProtected(), handler.UserList)
	app.Get("/user/:id", middleware.JWTProtected(), handler.UserDetail)
	app.Put("/user/:id", middleware.JWTProtected(), handler.UserEdit)
}

func (handler *UserController) UserList(c *fiber.Ctx) error {
	responses, err := handler.UserService.List()
	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Get Data Success",
		Error:   nil,
		Data:    responses,
	})
}

func (handler *UserController) UserDetail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.UserService.Detail(id)
	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Get Data Success",
		Error:   nil,
		Data:    response,
	})
}

func (handler *UserController) UserEdit(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	var input model.UpdateUserRequest
	err = c.BodyParser(&input)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.UserService.Edit(id, input)

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Update Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Update Data Success",
		Error:   nil,
		Data:    response,
	})
}
