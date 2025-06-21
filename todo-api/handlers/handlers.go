package handlers

import(
	"github.com/gofiber/fiber/v2"
	"fmt"
	"strconv"
)

type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}
var todos []Todo

func Gettodo(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "Hello and Welcome!"})
}

func Addtodo(c *fiber.Ctx) error {
	todo := Todo{}
	if err := c.BodyParser(&todo); err != nil  {
		return c.Status(500).JSON(fiber.Map{"Error": fmt.Sprint(err)})
	}
	if todo.Title == "" {
		return c.Status(400).JSON(fiber.Map{"body":"need to fill out the todo"})
	}
	todos = append(todos, todo)
	return c.Status(200).JSON(todos)
}

func Updatetodo(c *fiber.Ctx)error {
	id := c.Params("id")
	id_Int , err := strconv.Atoi(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Error":"Invalid ID"})
	}
	for i, todo := range todos{
		if todo.Id == id_Int { 
			todos[i].Done = !todo.Done
			return c.Status(200).JSON(todos[i])
		}
	}
	return c.Status(400).JSON(fiber.Map{"Message":"todo was not found"})
}

func Deletetodo(c *fiber.Ctx)error {
	id := c.Params("id")
	id_Int,err :=  strconv.Atoi(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Error":"Invalid ID"})	
	}
	for i, todo := range todos{
		if todo.Id == id_Int {
			todos = append(todos[:i], todos[i+1:]...) // the ... significe unpack the values
			return c.Status(200).JSON(fiber.Map{fmt.Sprint("Deletion of ", id):"is Succesful" })
		}
	}
	return c.Status(400).JSON(fiber.Map{"Message":"todo was not found"})	
}
