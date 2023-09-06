package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sraynitjsr/server/model"
)

func welcomeHandler(c *fiber.Ctx) error {
	return c.SendString("Welcome to Items Web Application")
}

func getItemHandler(c *fiber.Ctx, item model.Item) error {
	return c.JSON(item)
}

func SetupRoutes(app *fiber.App, item model.Item) {
	app.Get("/", welcomeHandler)

	app.Get("/getItemA", func(c *fiber.Ctx) error {
		itemA := item
		itemA.ID = 1
		itemA.Name = "Item A"
		itemA.Message = "Item A Fetched"
		return getItemHandler(c, itemA)
	})

	app.Get("/getItemB", func(c *fiber.Ctx) error {
		itemB := item
		itemB.ID = 2
		itemB.Name = "Item B"
		itemB.Message = "Item B Fetched"
		return getItemHandler(c, itemB)
	})
}
