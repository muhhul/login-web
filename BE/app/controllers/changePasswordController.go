package controllers

import "github.com/gofiber/fiber/v2"

func (server *Server) changePassword(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	if len(data["password"]) < 8 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be at least 8 characters",
		})
	}

	server.User.Password = data["password"]
	if err := server.DB.Save(&server.User).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not update the password",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Password changed successfully",
	})
}
