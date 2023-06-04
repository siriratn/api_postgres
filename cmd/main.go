// cmd/main.go

package main

 
import (
    "github.com/siriratn/api_postgres/database"
	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html"
)


func main() {
    database.ConnectDb()

    engine := html.New("./views", ".html") 
    app := fiber.New(fiber.Config{
        Views: engine, // new config
    })



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