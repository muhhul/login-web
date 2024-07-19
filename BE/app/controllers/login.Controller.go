package controllers

import (
	"login-web/app/models"
	"login-web/app/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (server *Server) Login(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var user models.User
	server.DB.Where("email = ?", data["email"]).First(&user)
	if user.Email == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email does not exist",
		})
	}

	if user.Password != data["password"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user":    user,
	})
}
