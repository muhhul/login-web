package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", Home).Methods("GET")
}

func (server *Server) Setup(app *fiber.App) {
	app.Post("/api/register", server.Register)
}
