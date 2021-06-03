package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vukani-dev/improvement-roll-service/controllers"
)

func CategoryRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
	route.Post("", controllers.CreateTodo)
	route.Get("/info", controllers.GetTodo)
}