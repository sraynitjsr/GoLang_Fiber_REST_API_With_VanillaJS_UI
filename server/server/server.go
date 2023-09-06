package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sraynitjsr/server/handler"
)

func main() {
	app := fiber.New()

	app.Get("/", handler.Home)
	app.Get("/getItems", handler.GetItems)
	app.Get("/getItemByID", handler.GetItemByID)
	app.Get("/getItemByName", handler.GetItemByName)

	app.Listen(":3000")
}
