// cmd/routes.go

package main


import (
    "github.com/gofiber/fiber/v2"
	"github.com/siriratn/api_postgres/handlers"
)

func setupRoutes(app *fiber.App) {
    //app.Get("/", handlers.Home)
    app.Get("/", handlers.ListFacts)
    app.Get("/fact", handlers.NewFactView) // Add new route for new view
   // Add new route to show single Fact, given `:id`
    app.Get("/fact/:id", handlers.ShowFact)
	app.Post("/fact", handlers.CreateFact)
}

/*func setupRoutes(app *fiber.App) {

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Nidnoi!")
    })

}*/