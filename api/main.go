package main

import (
	"database/sql"
)

func main() {
	var db *sql.DB

	db = connectPSQL()
	//this function runs at the end of the scope
	defer db.Close()

	//write code for graphql schema
}
