package main

import (
	"./book"
	"./database"
	"fmt"
	"github.com/gofiber/cors"
	_ "github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func getFile(c *fiber.Ctx) {
	err := c.SendFile("books.db")

	if err != nil {
		c.Next(err) // Pass error to Fiber
	}
}

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/file", getFile)

	apiV1 := app.Group("/api/v1", cors.New())

	apiV1.Get("/book", book.GetBooks)
	apiV1.Get("/book/:id", book.GetBook)
	apiV1.Post("/book", book.NewBook)
	apiV1.Delete("/book/:id", book.DeleteBook)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)
	app.Listen(3000)
}
