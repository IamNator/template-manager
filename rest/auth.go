package rest

import (
	"github.com/gofiber/fiber/v2"

	"template-manager/dto"
)

func (s server) Signup(c *fiber.Ctx) error {
	ctx := c.Context()

	var request dto.SignUpRequest

	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	err = s.App.Signup(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "check your email to continue sign up",
		"status":  true,
	})
}

func login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}
