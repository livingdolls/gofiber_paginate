package main

import (
	"fmt"
	"gofiber-paginate/controllers"
	"gofiber-paginate/database"
	"gofiber-paginate/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.InitDB()

	// add routes
	app.Get("/books/seed", func(c *fiber.Ctx) error {
		var book models.Book

		if err := database.DB.Exec("delete from books where 1").Error; err != nil {
			return c.Status(500).JSON("unable to delete data")
		}

		for i := 1; i <= 20; i++ {
			book.Title = fmt.Sprintf("Book %d", i)
			book.Description = fmt.Sprintf("This is a description for a book %d", i)
			book.Price = 500
			book.Author = fmt.Sprintf("Book author %d", i)
			book.CreatedAt = time.Now().Add(-time.Duration(21 - i) * time.Hour)
			database.DB.Create(&book)	
		}
		
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/books", controllers.GetPaginatedBooks)

	log.Fatal(app.Listen(":3000"))
}