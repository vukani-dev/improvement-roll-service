package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

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