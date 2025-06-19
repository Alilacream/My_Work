package db 

import (

	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB	
func Connect(){
	var err error
//That’s called a DSN (Data Source Name), and yeah — for Go to connect to MySQL using database/sql,
//  this format is required by the driver you're using (github.com/go-sql-driver/mysql).v
// meaning: root:root== [username]:[password] the adress==[host]  3306:port and finally todo_DB is the database name
	dsn := "root:root@tcp(127.0.0.1:3306)/todo" 
	DB , err := sql.Open("mysql", db)
	if err != nil {
		log.Fatal("An error occured to the open of the database: ", err)
	}
	err = DB.Ping()
	if err != nil {
		log Fatal("An error occured, couldn't ping the database: ", err)
	}
	log.Println("The connection to the database is succesfully permited.")

}