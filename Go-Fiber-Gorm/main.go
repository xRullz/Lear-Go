package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/route"

	"github.com/gofiber/fiber/v2"
)

// import "fmt"

func main() {
	//INITIAL DB
	database.ConnectDB()
	
	app := fiber.New()

	//INITIAL ROUTE
	route.RouteInit(app)

	app.Listen(":8080")
}