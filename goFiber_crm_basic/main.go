package main

import (
	"fmt"
	"goFiber_crm_basic/database"
	"goFiber_crm_basic/lead"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLeads)
	app.Delete("/api/v1/lead:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db") // open or error

	if err != nil {
		panic("failed connection db")
	}
	fmt.Println("Connection opened to db")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Db Migrated")
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close() // defer set at end of func

}
