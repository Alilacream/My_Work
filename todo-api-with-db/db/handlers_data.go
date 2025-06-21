package handlers_data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)
type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}
func Gettodo(id int)error {
	_, err := DB.Exec("INSERT INTO username id = (?)", id)
	return err
}