package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vukani-dev/improvement-roll-service/config"
	"github.com/vukani-dev/improvement-roll-service/routes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)


func setupRoutes(app *fiber.App) {
	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	// api group
	api := app.Group("/api")

	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	// send todos route group to TodoRoutes of routes package
	routes.CategoryRoute(api.Group("/categories"))
}

func main() {
  app := fiber.New()
	app.Use(logger.New())
	
  // Load .env
  err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}


  app.Get("/:search?/:filter?", func(c *fiber.Ctx) error {
	searchString := c.Params("search")
	filterString := c.Params("filter")
    fmt.Println(searchString)
    fmt.Println(filterString)

    return c.SendString("Hello, Hero!")
  })

  app.Listen(":3000")
}