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
		return c.SendString("Welcome to Items Web Application")
	})

	app.Get("/getItemA", func(c *fiber.Ctx) error {
		data := map[string]string{
			"message": "Item A Fetched",
		}
		return c.JSON(data)
	})

	app.Get("/getItemB", func(c *fiber.Ctx) error {
		data := map[string]string{
			"message": "Item B Fetched",
		}
		return c.JSON(data)
	})

	log.Fatal(app.Listen(":3000"))
}
