package handlers


import (
    "github.com/gofiber/fiber/v2"
	"github.com/siriratn/api_postgres/database"
	"github.com/siriratn/api_postgres/models"


)


func Hometest(c *fiber.Ctx) error {
    return c.SendString("Helllo Nidnoi :)")
}


func Home(c *fiber.Ctx) error {
    facts := []models.Fact{}
    database.DB.Db.Find(&facts)

    return c.Render("index", fiber.Map{
        "Title": "Many Time",
        "Subtitle": "Facts for funtimes with friends!",
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