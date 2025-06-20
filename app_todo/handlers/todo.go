package handlers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"todo/db"
)

// Handler for GET /todo
func Handler(c *fiber.Ctx) error {
	todos, err := db.GetTodos()
	if err != nil {
		return c.Status(500).SendString("failed to get todos\n")
	}
	return c.JSON(todos)
}

type Requirement struct {
	Purpose string `json:"purpose"`
}

type Completed struct {
	Done bool `json:"done"`
}

// Handler for POST /todo
func AddTodo(ctx *fiber.Ctx) error {
	var req Requirement

	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(400).SendString("Invalid input")
	}
	if req.Purpose == "" {
		return ctx.Status(400).SendString("Todo empty.")
	}
	err = db.AddTodo(req.Purpose)
	if err != nil {
		return ctx.Status(500).SendString("Invalid Todo: couldn't be added")
	}
	return ctx.Status(200).SendString("Added Todo list")
}

// Handler for DELETE /todo/:id
func DeleteTodo(ctx *fiber.Ctx) error {
	String_id := ctx.Param("id")
	Number_id, err := strconv.Atoi(String_id)
	if err != nil {
		return ctx.Status(500).SendString("Invalid input: not a number ID")
	}
	if err := db.DeleteTodo(Number_id); err != nil {
		return ctx.Status(400).SendString("Id does not exist")
	}
	return ctx.JSON(fiber.Map{"message": "The todo is fully deleted successfully."})
}

// Handler for PATCH /todo/:id
func Update(ctx *fiber.Ctx) error {
	String_id := ctx.Param("id")
	Number_id, err := strconv.Atoi(String_id)
	if err != nil {
		return ctx.Status(500).SendString("Invalid input: not a number ID")
	}
	var req Completed
	err = ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(400).SendString("Failed to parse Body")
	}
	if err := db.UpdateTodo(Number_id, req.Done); err != nil {
		return ctx.Status(500).SendString("Failed to update todo")
	}
	return ctx.SendString("Fully updated the todo")
}