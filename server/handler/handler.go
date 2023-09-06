package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sraynitjsr/server/model"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Welcome to Items API Home Page")
}

func GetItems(c *fiber.Ctx) error {
	items := []model.ItemStruct{
		{ID: 1, Name: "Item 1", Message: "Message One"},
		{ID: 2, Name: "Item 2", Message: "Message Two"},
		{ID: 3, Name: "Item 3", Message: "Message Three"},
		{ID: 4, Name: "Item 4", Message: "Message Four"},
		{ID: 5, Name: "Item 5", Message: "Message Five"},
	}

	return c.JSON(items)
}
