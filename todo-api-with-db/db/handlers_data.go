package handlers_data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var Todos []Todo
var DB *sql.DB // Make sure to assign this from your db package

func InsertTodo(title string) error {
	_, err := DB.Exec("INSERT INTO todos (title) VALUES (?)", title)
	return err
}

func UpdateTodo(id int, title string, done bool) error {
	_, err := DB.Exec("UPDATE todos SET title = ?, done = ? WHERE id = ?", title, done, id)
	return err
}

func DeleteTodo(id int) error {
	_, err := DB.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

func GetTodo(id int) (Todo, error) {
	row := DB.QueryRow("SELECT id, title, done FROM todos WHERE id = ?", id)
	var todo Todo
	err := row.Scan(&todo.Id, &todo.Title, &todo.Done)
	return todo, err
}

func GetAllTodos() ([]Todo, error) {
	rows, err := DB.Query("SELECT id, title, done FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}