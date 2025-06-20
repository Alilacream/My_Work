package main

import(
	"todo/handlers"
	"github.com/gin-gonic/gin"
	"todo/db"
)

func main() {
	db.Connect()
	app:= gin.Default()

	app.GET("/todo", handlers.Handler)
	app.POST("/todo", handlers.AddTodo)
	app.DELETE("/todo/:id", handlers.DeleteTodo)
	app.PATCH("/todo/:id", handlers.Update)

	app.Run(":3000")
}