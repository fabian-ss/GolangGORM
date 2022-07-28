package main

import (
	"github.com/fabian-ss/GolangGORM/app/routes"
	_ "github.com/fabian-ss/GolangGORM/docs"
	"github.com/gofiber/fiber/v2"
)

// @title Api usando Fiber, Swagger y GORM
// @version 1.0.0
// @description Api basica usando GORM.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	app := fiber.New()
	routes.SwaggerRoute(app)
	routes.RutaSaludo(app)
	app.Listen(":8080")
}
