package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sraynitjsr/server/db"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Welcome to Items API")
}

func GetItems(c *fiber.Ctx) error {
	items := db.GetAllItems()
	return c.JSON(items)
}

// http://localhost:3000/getItemByID?id=3 for example

func GetItemByID(c *fiber.Ctx) error {
	idStr := c.Query("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid 'id' parameter")
	}

	item, found := db.GetItemByID(id)
	if !found {
		return c.Status(fiber.StatusNotFound).SendString("Item not found")
	}

	return c.JSON(item)
}

// http://localhost:3000/getItemByName?name=Item%203 for example

func GetItemByName(c *fiber.Ctx) error {
	name := c.Query("name")

	item, found := db.GetItemByName(name)
	if !found {
		return c.Status(fiber.StatusNotFound).SendString("Item not found")
	}

	return c.JSON(item)
}
