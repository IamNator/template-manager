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

	err = s.app.Signup(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "check your email to continue sign up",
		"status":  true,
	})
}

func (s server) Login(c *fiber.Ctx) error {
	var request dto.LoginRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	response, err := s.app.Login(c.Context(), request)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "welcome back",
		"body":    response,
	})
}

func (s server) Logout(c *fiber.Ctx) error {
	var request dto.LogoutRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	err = s.app.Logout(c.Context(), request)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "successfully logged out",
	})
}
