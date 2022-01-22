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

type AssignController struct {
	AssignService service.AssignService
}

func NewAssignController(assignService *service.AssignService) AssignController {
	return AssignController{*assignService}
}

func (handler *AssignController) Route(app *fiber.App) {
	app.Post("/assign/user-role", middleware.JWTProtected(), handler.SetUserRole)
	app.Post("/assign/role-permission", middleware.JWTProtected(), handler.SetRolePermission)
}

func (handler *AssignController) SetUserRole(c *fiber.Ctx) error {
	var input model.UserRoleRequest

	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.AssignService.UserRole(input)

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Assign Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Assign Data Success",
		Error:   nil,
		Data:    response,
	})
}

func (handler *AssignController) SetRolePermission(c *fiber.Ctx) error {
	var input model.RolePermissionRequest

	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.AssignService.RolePermission(input)

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Assign Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Assign Data Success",
		Error:   nil,
		Data:    response,
	})
}
