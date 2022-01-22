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

type PermissionController struct {
	PermissionService service.PermissionService
}

func NewPermissionController(permissionService *service.PermissionService) PermissionController {
	return PermissionController{*permissionService}
}

func (handler *PermissionController) Route(app *fiber.App) {
	app.Get("/permission", middleware.JWTProtected(), handler.PermissionList)
	app.Get("/permission/:id", middleware.JWTProtected(), handler.PermissionDetail)
	app.Post("/permission", middleware.JWTProtected(), handler.PermissionCreate)
	app.Put("/permission/:id", middleware.JWTProtected(), handler.PermissionEdit)
	app.Delete("/permission/:id", middleware.JWTProtected(), handler.PermissionDelete)
}

func (handler *PermissionController) PermissionList(c *fiber.Ctx) error {
	responses, err := handler.PermissionService.List()
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

func (handler *PermissionController) PermissionDetail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.PermissionService.Detail(id)
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

func (handler *PermissionController) PermissionCreate(c *fiber.Ctx) error {
	var input model.CreatePermissionRequest
	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.PermissionService.Create(input)

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

func (handler *PermissionController) PermissionEdit(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	var input model.UpdatePermissionRequest
	err = c.BodyParser(&input)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.PermissionService.Edit(id, input)

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

func (handler *PermissionController) PermissionDelete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	err = handler.PermissionService.Delete(id)

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
