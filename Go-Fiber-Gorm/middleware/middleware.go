package middleware

import "github.com/gofiber/fiber/v2"

func Auth(ctx *fiber.Ctx) error {
	auth := ctx.Get("x-token")
	if auth != "secret" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	return ctx.Next()
}

func Permission(ctx *fiber.Ctx) error {
	
	return ctx.Next()
}