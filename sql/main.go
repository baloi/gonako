package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

func main() {
	fmt.Println("vim-go")
	db, _ := sql.Open("sqlite3", "./baloi.db")
	defer db.Close()

	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS people(id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = db.Prepare("INSERT into people(firstname, lastname) VALUES (?, ?)")
	statement.Exec("Nigel", "Gali")
	rows, _ := db.Query("SELECT id, firstname, lastname from people")

	var id int
	var firstname string
	var lastname string

	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}

	rows.Close()

}
