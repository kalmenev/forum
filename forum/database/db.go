package database

import (
	"database/sql"
	"log"

	"forum/config"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("sqlite3", config.DATABASE)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if _, err := DB.Exec(Schema); err != nil {
		log.Println(err.Error())
		return
	}

	
}
