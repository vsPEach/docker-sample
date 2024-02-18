package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]string{
			"status": "OK",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
