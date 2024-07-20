package controllers

import (
	"login-web/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func (server *Server) Setup(app *fiber.App) {
	app.Post("/api/register", server.Register)
	app.Post("/api/login", server.Login)
	app.Post("/api/changePassword", server.changePassword)
	app.Post("/api/deleteAccount", server.deleteAccount)
	app.Use(middleware.IsAuthenticate)
}
