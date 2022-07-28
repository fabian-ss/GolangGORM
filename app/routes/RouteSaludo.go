package routes

import (
	"github.com/fabian-ss/GolangGORM/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func RutaSaludo(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api")

	// Routes for GET method:
	route.Get("/saludo", controllers.Saludo)

}
