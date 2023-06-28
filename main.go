package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Tufan"},
	{ID: 2, Name: "CR7"},
}

func rootHandler(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Users REST API Using Go Fiber")
}

func getUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func getUserById(c *fiber.Ctx) error {
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

func deleteUserById(c *fiber.Ctx) error {
	return c.SendString("To be implemented yet")
}

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Routes
	app.Get("/", rootHandler)
	app.Get("/users", getUsers)
	app.Get("/users/id", getUserById)
	app.Delete("/users/id", deleteUserById)

	// Create a channel to listen for an interrupt or termination signal from the OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Start the server in a separate goroutine
	go func() {
		if err := app.Listen(":8000"); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for the interrupt signal
	<-quit
	log.Println("Shutting down server...")

	// Create a context with a timeout for the graceful shutdown
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Perform the graceful shutdown by shutting down the server and blocking until all connections are closed
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Graceful shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped")
}
