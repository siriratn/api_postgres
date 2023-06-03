// cmd/routes.go

package main


import (
    "github.com/gofiber/fiber/v2"
	"github.com/siriratn/api_postgres/handlers"
)

func setupRoutes(app *fiber.App) {
    app.Get("/", handlers.Home)
}

func setupRoutes(app *fiber.App) {

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Nidnoi!")
    })

}