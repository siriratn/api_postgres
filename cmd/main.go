// cmd/main.go

package main


import (
    "github.com/siriratn/api_postgres/database"
    "github.com/siriratn/api_postgres/handlers"
	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html"
)


func main() {
    database.ConnectDb()

    engine := html.New("./views", ".html") 

    app := fiber.New(fiber.Config{
        Views: engine,
        ViewsLayout: "layouts/main", // add this to config
    })


    setupRoutes(app)
    
    app.Static("/", "./public")

    // app.Get("/", func(c *fiber.Ctx) error {
    //     return c.SendString("Hello, Nidnoi!")
    // })

    // Set up 404 page
    app.Use(handlers.NotFound) //Remember to place this after the other setup, but right before the call to app.Listen

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