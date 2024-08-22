package route

import (
	"go-fiber-gorm/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/users", controller.GetAll)
	r.Get("/users/:id", controller.Detail)
	r.Post("/users", controller.Create)
	r.Put("/users/:id", controller.Update)
	r.Delete("/users/:id", controller.Delete)
}
