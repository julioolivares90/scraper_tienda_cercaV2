package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/julioolivares90/scraper_tienda_cercaV2/scraper"
)

func FindLugar(c *fiber.Ctx) {
	lugar := c.Query("lugar")

	response := scraper.FindLugar(lugar)

	c.JSON(response)
}
