package router

import (
	"github.com/edr3x/fiber-starter/modules/user"
	"github.com/gofiber/fiber/v2"
)

func MainRouter(r fiber.Router) {
	user.Routes(r.Group("/user"))

	// Add your routes here like this:
	// auth.Routes(r.Group("/auth"))
	// post.Routes(r.Group("/post", middleware.RequireAuth())) // can call middleware here
}
