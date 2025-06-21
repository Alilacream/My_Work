package main

import (
	"Todo/handlers"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	fmt.Println("Welcome to the fiber server")
	app := fiber.New()

	// Get Todo
	app.Get("/", handlers.Gettodo)
	// Create Todo
	app.Post("/todo", handlers.Addtodo)
	// Update Todo
	app.Patch("/todo/:id", handlers.Updatetodo)
	// Delete Todo
	 app.Delete("/todo/:id", handlers.Deletetodo)
	log.Fatal(app.Listen(":" + PORT))
}