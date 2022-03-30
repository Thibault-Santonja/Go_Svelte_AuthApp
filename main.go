package main

import (
	auth "github.com/Thibault-Santonja/Go_Svelte_AuthApp/server"
	"github.com/gofiber/fiber/v2"
)

func launchAppListener(app *fiber.App) {
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}

func initRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	auth.SetAuthRoutes(app)
}

func main() {
	app := fiber.New()
	initRoutes(app)
	launchAppListener(app)
}
