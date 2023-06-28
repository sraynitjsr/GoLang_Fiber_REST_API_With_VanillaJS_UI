package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Subhradeep Ray"},
	{ID: 2, Name: "Sachin Tendulkar"},
}

func home(c *fiber.Ctx) error {
	welcome := "Welcome to Users REST API"
	return c.SendString(welcome)
}

func getUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error {
	id := c.Params("id")

	for _, user := range users {
		if fmt.Sprint(user.ID) == id {
			return c.JSON(user)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "User not found",
	})
}

func main() {
	app := fiber.New()

	app.Get("/", home)
	app.Get("/users", getUsers)
	app.Get("/users/:id", getUser)

	log.Fatal(app.Listen(":8000"))
}
