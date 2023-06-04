package handlers


import (
    "github.com/gofiber/fiber/v2"
	"github.com/siriratn/api_postgres/database"
	"github.com/siriratn/api_postgres/models"

)


func Hometest(c *fiber.Ctx) error {
    return c.SendString("Helllo Nidnoi :)")
}


func ListFacts(c *fiber.Ctx) error {
    facts := []models.Fact{}
    database.DB.Db.Find(&facts)

    return c.Render("index", fiber.Map{
        "Title": "Many Time",
        "Subtitle": "Facts for funtimes with friends!",
       "Facts":    facts, // send the facts to the view
    })
    
}

// Create new Fact View handler
func NewFactView(c *fiber.Ctx) error {
    return c.Render("new", fiber.Map{
        "Title":    "New Fact",
        "Subtitle": "Add a cool fact!",
    })
}

func CreateFact(c *fiber.Ctx) error {
    fact := new(models.Fact)
    if err := c.BodyParser(fact); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    database.DB.Db.Create(&fact)

    return c.Status(200).JSON(fact)
}