package db

import (
	"fmt"
	"database/sql"
	"os"
	"log"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't load the Env File: ",err )
	}
	dsn := os.Getenv("MYSQL_DSN")
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Data Source Name inrecognizable: ", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("Couldn't ping the Database: ", err)
	}
	fmt.Println("Succesfully entered the Database✅")
}