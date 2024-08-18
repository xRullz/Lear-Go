package controller

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/models/entity"
	"go-fiber-gorm/models/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)

	// return nil
}

func Create(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	newUser := entity.User{
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error

	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success",
		"data": newUser,
	})
}
