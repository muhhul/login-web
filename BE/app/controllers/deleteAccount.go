package controllers

import "github.com/gofiber/fiber/v2"

func (server *Server) deleteAccount(c *fiber.Ctx) error {
	if err := server.DB.Delete(&server.User).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete account",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Account deleted successfully",
	})
}
