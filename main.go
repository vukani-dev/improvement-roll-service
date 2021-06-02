package main

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func main() {
  app := fiber.New()

  app.Post("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, Hero!")
  })

  app.Get("/:search?/:filter?", func(c *fiber.Ctx) error {
	searchString := c.Params("search")
	filterString := c.Params("filter")
    fmt.Println(searchString)
    fmt.Println(filterString)

    return c.SendString("Hello, Hero!")
  })

  app.Get("/:value", func(c *fiber.Ctx) error {
  return c.SendString("value: " + c.Params("value"))
  // => Get request with value: hello world
})
  app.Listen(":3000")
}