package auth

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
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
	app.Get("/auth", getUsers)
	app.Get("/auth/:value", getUser)
}

func getUsers(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(people)
}

func getUser(c *fiber.Ctx) error {
	valueStr := c.Params("value")
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil || valueInt >= len(people) {
		return c.Status(fiber.StatusNotFound).JSON("This ID is not found")
	}
	return c.Status(fiber.StatusOK).JSON(people[valueInt])
}
