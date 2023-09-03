package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/helloWorld", func(c *fiber.Ctx) error {
		data := map[string]string{
			"message": "This is a Sample API Endpoint",
		}
		return c.JSON(data)
	})

	log.Fatal(app.Listen(":3000"))
}
