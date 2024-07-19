package controllers

import (
	"log"
	"login-web/app/models"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return Re.MatchString(email)
}

func (server *Server) Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid email",
		})
	}

	server.DB.Where("email = ?", data["email"].(string)).First(&userData)
	if userData.Email != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already registered",
		})
	}

	if len(data["password"].(string)) < 8 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be at least 8 characters",
		})
	}

	user := models.User{
		Email:    data["email"].(string),
		Password: data["password"].(string),
	}

	result := server.DB.Create(&user)
	if result.Error != nil {
		log.Println("3")
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Could not register user",
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "User registered successfully",
	})
}
