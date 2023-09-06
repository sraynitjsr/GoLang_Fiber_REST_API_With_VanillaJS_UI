package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sraynitjsr/server/config"
	"github.com/sraynitjsr/server/db"
	"github.com/sraynitjsr/server/handler"
)

func main() {
	appConfig, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	}

	app := fiber.New()

	itemService := db.NewItemDB()

	itemHandler := handler.NewItemHandler(itemService)

	app.Get("/", itemHandler.Home)
	app.Get("/getItems", itemHandler.GetItems)
	app.Get("/getItemByID", itemHandler.GetItemByID)
	app.Get("/getItemByName", itemHandler.GetItemByName)

	port := fmt.Sprintf(":%d", appConfig.Port)
	app.Listen(port)
}
