package db 

import(
	"database/sql"
	"log"
)
type Todo struct {
	ID int `json:"id"`
	Purpose string `json:"purpose`
	Done bool `json:""done`
}
// Create CRUD
func AddTodo(purpose string)error {
	
	_,err := DB.Exec("INSERT INTO todos (purpose) VALUES (?)", purpose)
	return err
}
// Read
func GetTodos() ([]Todo,error) {
	rows,err:= DB.Query("SELECT id, purpose, done FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []Todo
	for rows.Next() {// the reason why rows.Next() only it is because that functions returns a boolean not a value
		var t Todo
		err := rows.Scan(&t.ID, &t.Purpose, &t.Done) 
		if err != nil {
		return nil , err
		}
	}
	todos = append(todos, t)
	return todos, nil
}
// delete CRUD
func DeleteTodo(id int) error {
	_, err := DB.Exec("DELETE FROM todos WHERE id == ?", id)
	return err
} 

func UpdateTodo(id int, done bool) {
	_, err:= DB.Exec("UPDATE todos SET done == ? WHERE id == ?", done,id)
	return err
}
