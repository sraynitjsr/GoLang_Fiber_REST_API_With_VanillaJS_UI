package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sraynitjsr/server/model"
)

func welcomeHandler(c *fiber.Ctx) error {
	return c.SendString("Welcome to Items Web Application")
}

func getItemA(c *fiber.Ctx) error {
	item := model.ItemStruct{
		ID:      1,
		Name:    "A",
		Message: "This is A",
	}
	return c.JSON(item)
}

func getItemB(c *fiber.Ctx) error {
	item := model.ItemStruct{
		ID:      2,
		Name:    "B",
		Message: "This is B",
	}
	return c.JSON(item)
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", welcomeHandler)
	app.Get("/getItemA", getItemA)
	app.Get("/getItemB", getItemB)
}
