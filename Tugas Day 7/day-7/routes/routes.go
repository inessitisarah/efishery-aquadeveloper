package routes

import (
	"mymodule/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router, userHandler *handler.UserHandler) {
	r := app.Group("api/v1")

	r.Get("/users", userHandler.GetList)
	r.Post("/user", userHandler.CreateUser)
	r.Delete("/admin/:id", userHandler.DeleteUser)

}
