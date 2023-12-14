package rest

import "github.com/gofiber/fiber/v2"

func addTemplate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func getTemplate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func deleteTemplate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func updateTemplate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}
