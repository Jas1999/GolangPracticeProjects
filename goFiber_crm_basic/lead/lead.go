package lead

import (
	"goFiber_crm_basic/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

// `json:"name"` telling wht it will look like in json

// context is being passed with api path
// func can work with data given from context
func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead // slice is array of object
	db.Find(&leads)
	c.JSON(leads)

}
func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.Find(&lead, id) // response being sent via lead
	c.JSON(lead)
}
func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead) // new instance
	// body parser checks that data matches lead first

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)

}
func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead deleted")
}
