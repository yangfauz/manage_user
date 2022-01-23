package main

import (
	"net/http"
	"service-acl/config"
	"service-acl/controller"
	"service-acl/entity/migration"
	"service-acl/model/responder"
	"service-acl/repository"
	"service-acl/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Setup configuration
	configuration := config.New()
	database := config.NewMysqlDatabase(configuration)

	// Setup Migration
	migration.Migration(database)

	// Setup Repository
	userRepository := repository.NewUserRepository(database)
	roleRepository := repository.NewRoleRepository(database)
	permissionRepository := repository.NewPermissionRepository(database)
	menuRepository := repository.NewMenuRepository(database)

	// Setup Service
	authService := service.NewAuthService(&userRepository)
	userService := service.NewUserService(&userRepository)
	profileService := service.NewProfileService(&userRepository, &menuRepository)
	assignService := service.NewAssignService(&userRepository, &roleRepository)
	roleService := service.NewRoleService(&roleRepository)
	permissionService := service.NewPermissionService(&permissionRepository)
	menuService := service.NewMenuService(&menuRepository)

	// Setup Controller
	authController := controller.NewAuthController(&authService)
	userController := controller.NewUserController(&userService)
	profileController := controller.NewProfileController(&profileService)
	assignController := controller.NewAssignController(&assignService)
	roleController := controller.NewRoleController(&roleService)
	permissionController := controller.NewPermissionController(&permissionService)
	menuController := controller.NewMenuController(&menuService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(responder.ApiResponse{
			Code:    http.StatusOK,
			Message: "User Management Service",
			Error:   nil,
			Data:    nil,
		})
	})

	// Setup Routing
	authController.Route(app)
	userController.Route(app)
	profileController.Route(app)
	assignController.Route(app)
	roleController.Route(app)
	permissionController.Route(app)
	menuController.Route(app)

	//Not Found in Last
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:  http.StatusNotFound,
			Error: &fiber.ErrNotFound.Message,
			Data:  nil,
		})
	})

	// Start App
	port := configuration.Get("APP_PORT")
	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
