package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sraynitjsr/server/handler"
)

func main() {
	app := fiber.New()

	app.Get("/", handler.Home)
	app.Get("/getItems", handler.GetItems)

	app.Listen(":3000")
}
