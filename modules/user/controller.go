package user

import (
	"github.com/edr3x/fiber-starter/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUserController(c *fiber.Ctx) error {
	return c.JSON(utils.SuccessResponse{
		Success: true,
		Payload: "Hello From User",
	})
}

func GetUserByIdController(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(utils.SuccessResponse{
		Success: true,
		Payload: "Hello From User " + id,
	})
}
