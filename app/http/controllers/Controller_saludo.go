package controllers

import "github.com/gofiber/fiber/v2"

// @Summary Es una función que devuelve un saludo
// @Description obten un saludo
// @Accept  json
// @Produce  json
// @Success 200 string 200
// @Router /api/saludo [get]
// @Tags Saludo
func Saludo(c *fiber.Ctx) error {
	return c.SendString("Hola mundo")
}
