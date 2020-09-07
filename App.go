package main

import (
	"os"

	"github.com/gofiber/fiber"
	"github.com/julioolivares90/scraper_tienda_cercaV2/controllers"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("hello tienda-cerca")
	})

	app.Get("/api/find-lugar", controllers.FindLugar)
	port := getPort()
	app.Listen(port)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}
	return port
}
