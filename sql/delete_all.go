package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	//"strconv"
)

func main() {
	fmt.Println("Will delete all data from database")

	db, _ := sql.Open("sqlite3", "./baloi.db")
	defer db.Close()

	statement, _ := db.Prepare("DELETE FROM PEOPLE")
	statement.Exec()

	rows, _ := db.Query("SELECT * FROM PEOPLE")
	statement.Exec()

	var id int
	var firstname string
	var lastname string
	var hasRows = false

	for rows.Next() {
		hasRows = true
		if err := rows.Scan(&id, &firstname, &lastname); err != nil {
			fmt.Println("error")
		} else {
			fmt.Println("no error")
		}

		//fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}

	if hasRows {
		fmt.Println("Has rows")
	} else {
		fmt.Println("No rows")
	}

	rows.Close()
}
