package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/stheven26/config"
	"github.com/stheven26/db"
	"github.com/stheven26/globals"
	"github.com/stheven26/helpers"
	"github.com/stheven26/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"time":    time.Now(),
		})
	}

	hash, err := helpers.HashPassword(data["password"])

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"messgae": err.Error(),
			"time":    time.Now(),
		})
	}

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: hash,
	}

	connection := db.GetConnection()
	connection.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	var user models.User

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	connection := db.GetConnection()

	connection.Where(`email = ?`, data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
			"time":    time.Now(),
		})
	}

	match, err := helpers.CheckHashPassword(data["password"], user.Password)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(fiber.Map{
			"message": "password incorrect!",
			"time":    time.Now(),
		})
	}

	if !match {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"time":    time.Now(),
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
	})

	token, err := claims.SignedString([]byte(globals.Key))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"time":    time.Now(),
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(72 * time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Success Login!",
		"time":    time.Now(),
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Success Logout!",
		"time":    time.Now(),
	})
}

func GetData(c *fiber.Ctx) error {
	var data models.Data
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(globals.Key), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
			"time":    time.Now(),
		})
	}

	decodeData := config.ParsingJson()

	for _, v := range decodeData.Data {
		data.ID = v.ID
		for _, v2 := range v.Details {
			data.Name = v2.Name
			data.Balance = v2.Balance
		}
		data.Transportation = v.FavoriteTransportation
	}

	connection := db.GetConnection()
	connection.Create(&data)

	connection.Where("balance <= ?, 10000").Find(&data)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"data": data,
	})
}
