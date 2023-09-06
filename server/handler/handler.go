package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sraynitjsr/server/service"
)

type ItemHandler struct {
	ItemService service.ItemService
}

func NewItemHandler(itemService service.ItemService) *ItemHandler {
	return &ItemHandler{
		ItemService: itemService,
	}
}

// curl http://localhost:3000/getItemByID?id=1

func (h *ItemHandler) GetItemByID(c *fiber.Ctx) error {
	idStr := c.Query("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid 'id' parameter")
	}

	item, found := h.ItemService.GetItemByID(id)
	if !found {
		return c.Status(fiber.StatusNotFound).SendString("Item not found")
	}

	return c.JSON(item)
}

// curl http://localhost:3000/getItemByName?name=Item%203

func (h *ItemHandler) GetItemByName(c *fiber.Ctx) error {
	name := c.Query("name")

	item, found := h.ItemService.GetItemByName(name)
	if !found {
		return c.Status(fiber.StatusNotFound).SendString("Item not found")
	}

	return c.JSON(item)
}

// curl http://localhost:3000/getItems

func (h *ItemHandler) GetItems(c *fiber.Ctx) error {
	items := h.ItemService.GetAllItems()
	return c.JSON(items)
}

func (h *ItemHandler) Home(c *fiber.Ctx) error {
	return c.SendString("Welcome to API For Items")
}
