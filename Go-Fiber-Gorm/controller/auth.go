package controller

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/helpers"
	"go-fiber-gorm/models/entity"
	"go-fiber-gorm/models/request"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	login := new(request.LoginRequest)
	if err := ctx.BodyParser(login); err != nil {
		return err
	}

	validate := validator.New()
	errvalidate := validate.Struct(login)
	if errvalidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed",
			"error":   errvalidate.Error(),
		})
	}

	//Check User
	var user entity.User
	err := database.DB.First(&user, "email = ?", login.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	//Validation Password
	isValid := helpers.CheckHashingPass(login.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	//Generate Token
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	token, errGenerate := helpers.GenerateToken(&claims)
	if errGenerate != nil {
		log.Println(errGenerate)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
