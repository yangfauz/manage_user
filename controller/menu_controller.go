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

type MenuController struct {
	MenuService service.MenuService
}

func NewMenuController(menuService *service.MenuService) MenuController {
	return MenuController{*menuService}
}

func (handler *MenuController) Route(app *fiber.App) {
	app.Get("/menu", middleware.JWTProtected(), handler.MenuList)
	app.Get("/menu/:id", middleware.JWTProtected(), handler.MenuDetail)
	app.Post("/menu", middleware.JWTProtected(), handler.MenuCreate)
	app.Put("/menu/:id", middleware.JWTProtected(), handler.MenuEdit)
	app.Delete("/menu/:id", middleware.JWTProtected(), handler.MenuDelete)
}

func (handler *MenuController) MenuList(c *fiber.Ctx) error {
	responses, err := handler.MenuService.List()
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

func (handler *MenuController) MenuDetail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.MenuService.Detail(id)
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

func (handler *MenuController) MenuCreate(c *fiber.Ctx) error {
	var input model.CreateMenuRequest
	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.MenuService.Create(input)

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

func (handler *MenuController) MenuEdit(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	var input model.UpdateMenuRequest
	err = c.BodyParser(&input)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	response, err := handler.MenuService.Edit(id, input)

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

func (handler *MenuController) MenuDelete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	err = handler.MenuService.Delete(id)

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
