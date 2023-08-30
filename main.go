package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendFile("./limit-reached.html")
		},
		Storage: myCustomStorage{},
	}))

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Post("/login", login)

	app.Listen(3000)
}

func login(c *fiber.Ctx) error {
    user := c.FormValue("user")
    pass := c.FormValue("pass")

    if user != "ray" || pass != "007" {
        return c.SendStatus(fiber.StatusUnauthorized)
    }

    claims := jwt.MapClaims{
        "name":  "Subhradeep Ray",
        "admin": true,
        "exp":   time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    t, err := token.SignedString([]byte("secret"))
    if err != nil {
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    return c.JSON(fiber.Map{"token": t})
}
