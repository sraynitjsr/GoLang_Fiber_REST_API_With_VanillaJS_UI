package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func welcomeHandler(c *fiber.Ctx) error {
	return c.SendString("Welcome to Items Web Application")
}

func getItemAHandler(c *fiber.Ctx) error {
	data := map[string]string{
		"message": "Item A Fetched",
	}
	return c.JSON(data)
}

func getItemBHandler(c *fiber.Ctx) error {
	data := map[string]string{
		"message": "Item B Fetched",
	}
	return c.JSON(data)
}

func setupRoutes(app *fiber.App) {
	app.Use(cors.New())

	app.Get("/", welcomeHandler)
	app.Get("/getItemA", getItemAHandler)
	app.Get("/getItemB", getItemBHandler)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
