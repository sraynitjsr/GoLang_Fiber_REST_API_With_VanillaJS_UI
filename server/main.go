package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sraynitjsr/server/handler"
)

func main() {
	app := fiber.New()

	handler.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
