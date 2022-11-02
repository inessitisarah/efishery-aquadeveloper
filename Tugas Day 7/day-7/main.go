package main

import (
	"mymodule/config"
	"mymodule/handler"
	"mymodule/repository"
	"mymodule/routes"
	"mymodule/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Database()
	config.AutoMigrate()

	app := fiber.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	routes.Routes(app, userHandler)

	app.Listen(":3000")

	//app := fiber.New()
}
