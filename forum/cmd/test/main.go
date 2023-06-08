package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite3 driver
)

func main() {
	// Establish a database connection
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to ping the database:", err)
		return
	}

	// Prepare the SQL query
	query := `SELECT "user_id", "name", "session", "expire" from Sessions where "name"=?;`

	// Execute the query and retrieve the session value
	var userId, name, session string
	// var expires int64
	err = db.QueryRow(query, "2").Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No session found for the user: ", err.Error())
		} else {
			fmt.Println("Failed to query the session:", err.Error())
		}
		return
	}

	// Process the session value
	fmt.Println("user_id:", userId,"\nName:", name, "\nSession:", session)
}
