package auth

import (
	"github.com/gofiber/fiber/v2"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

var people = []Person{
	{ID: 0, Name: "admin", Pass: "pass"},
	{ID: 1, Name: "user", Pass: "password"},
}

func SetAuthRoutes(app *fiber.App) {
	app.Get("/auth", getAuth)
}

func getAuth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(people)
}
