package user

import "github.com/gofiber/fiber/v2"

func Routes(r fiber.Router) {
	r.Get("/", GetUserController)

	r.Get("/:id", GetUserByIdController)
}
