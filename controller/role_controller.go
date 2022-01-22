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

type RoleController struct {
	RoleService service.RoleService
}

func NewRoleController(roleService *service.RoleService) RoleController {
	return RoleController{*roleService}
}

func (handler *RoleController) Route(app *fiber.App) {
	app.Get("/role", middleware.JWTProtected(), handler.RoleList)
	app.Get("/role/:id", middleware.JWTProtected(), handler.RoleDetail)
	app.Post("/role", middleware.JWTProtected(), handler.RoleCreate)
	app.Put("/role/:id", middleware.JWTProtected(), handler.RoleEdit)
	app.Delete("/role/:id", middleware.JWTProtected(), handler.RoleDelete)
}

func (handler *RoleController) RoleList(c *fiber.Ctx) error {
	responses, err := handler.RoleService.List()
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

func (handler *RoleController) RoleDetail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.RoleService.Detail(id)
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

func (handler *RoleController) RoleCreate(c *fiber.Ctx) error {
	var input model.CreateRoleRequest
	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.RoleService.Create(input)

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Create Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Create Data Success",
		Error:   nil,
		Data:    response,
	})
}

func (handler *RoleController) RoleEdit(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	var input model.UpdateRoleRequest
	err = c.BodyParser(&input)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.RoleService.Edit(id, input)

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

func (handler *RoleController) RoleDelete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	err = handler.RoleService.Delete(id)

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Delete Data Success",
		Error:   nil,
		Data:    nil,
	})
}
