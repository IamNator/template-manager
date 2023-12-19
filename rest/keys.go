package rest

import (
	"github.com/gofiber/fiber/v2"
	"template-manager/dto"
)

func (s server) AddKey(c *fiber.Ctx) error {
	ctx := c.Context()

	var request dto.CreateAccessKeyRequest
	err := c.BodyParser(&request)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	err = s.app.CreateAccessKey(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func (s server) ListAccessKeys(c *fiber.Ctx) error {
	ctx := c.Context()

	var request dto.ListAccessKeysRequest
	keys, err := s.app.ListAccessKeys(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "pong",
		"keys":    keys,
	})
}

func (s server) DeleteKey(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}
