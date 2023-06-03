// cmd/main.go

package main

import "github.com/gofiber/fiber/v2"

func main() {
    database.ConnectDb()

    app := fiber.New()

    setupRoutes(app)

    // app.Get("/", func(c *fiber.Ctx) error {
    //     return c.SendString("Hello, Nidnoi!")
    // })

    app.Listen(":3000")
}

/*
func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Nidnoi!")
    })

    app.Listen(":3000")
}*/