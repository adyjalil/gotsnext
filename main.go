package main

import (
	"fmt"
	"gotsnext/internal/helpers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {

	conf := fiber.Config{
		ServerHeader: "go fiber",
	}
	app := fiber.New(conf)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/n/:number", func(c *fiber.Ctx) error {
		number := c.Params("number", "0")
		numberInt, err := strconv.Atoi(number)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).SendString("cannot convert int")
		}

		n := helpers.IntToString(numberInt)
		return c.SendString(n)
	})

	app.Get("/s/:number", func(c *fiber.Ctx) error {
		number := c.Params("number", "0")
		n, err := helpers.StringToInt(number)
		if err != nil {
			return fiber.ErrBadRequest
		}

		return c.SendString(fmt.Sprintf("number value: %d", n))
	})

	app.Listen(":3000")
}
