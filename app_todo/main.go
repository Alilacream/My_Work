package main

import(
	"todo/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app:= fiber.New()

	app.Get("/todo", handlers.Handler)
	app.Post("/todo", handlers.AddTodo)
	app.Delete("/todo/:id", handlers.DeleteTodo)
	app.Patch("/todo/:id", handlers.Update)

	app.Listen(":3000")
}