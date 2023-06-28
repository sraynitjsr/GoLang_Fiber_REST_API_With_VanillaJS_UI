package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func rootHandler(c *fiber.Ctx) error {
	return c.SendString("Welcome to the API!")
}

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Routes
	app.Get("/", rootHandler)

	// Create a channel to listen for an interrupt or termination signal from the OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Start the server in a separate goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for the interrupt signal
	<-quit
	log.Println("Shutting down server...")

	// Create a context with a timeout for the graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Perform the graceful shutdown by shutting down the server and blocking until all connections are closed
	if err := app.Shutdown(ctx); err != nil {
		log.Fatalf("Graceful shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped")
}
