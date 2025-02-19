package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:lzrwashere@/gocrud?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Databse Connected")
	DB = db
}
