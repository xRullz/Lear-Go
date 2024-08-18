package route

import (
	"go-fiber-gorm/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/users", controller.GetAll)
	r.Post("/users", controller.Create)
}
