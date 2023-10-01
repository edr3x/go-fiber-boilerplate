package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"

	router "github.com/edr3x/fiber-starter/modules"
	"github.com/edr3x/fiber-starter/utils"
)

type FailureResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Path    string `json:"url_path"`
}

func main() {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			var e *fiber.Error
			if ok := errors.As(err, &e); !ok {
				log.Println(err.Error())
				e = fiber.NewError(500, "Oops An Unexpected Error Occurred")
			}
			ctx.Status(e.Code).JSON(FailureResponse{
				Success: false,
				Message: e.Message,
				Path:    ctx.OriginalURL(),
			})
			return nil
		},
	})

	// Middlewares
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Security Headers
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("Strict-Transport-Security", "max-age=5184000; includeSubDomains")
		return c.Next()
	})

	// Secure Api
	// app.Use(func(c *fiber.Ctx) error {
	// 	if c.Get("x-access-key") != os.Getenv("ACCESS_KEY") {
	// 		return fiber.NewError(401, "Unauthorized")
	// 	}
	// 	return c.Next()
	// })

	// Rate Limiter
	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 30 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return fiber.NewError(429, "Too Many Requests")
		},
	}))

	// routecheck
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(utils.SuccessResponse{
			Success: true,
			Payload: "Hello There",
		})
	})

	// main router
	api := app.Group("/api/v1")
	router.MainRouter(api)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return fiber.NewError(404, "Endpoint Not Found")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
