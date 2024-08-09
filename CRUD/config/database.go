package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "rull:rull@/crud-go?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("berhasil konek")
	DB = db
}
