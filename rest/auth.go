package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"template-manager/dto"
	"template-manager/entity"
)

func signup(c *fiber.Ctx) error {
	var request dto.SignUpRequest

	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	var account = entity.Account{
		Email: request.Email,
	}

	// TODO: save to db

	fmt.Println(account)

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
