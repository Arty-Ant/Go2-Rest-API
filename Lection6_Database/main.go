package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"


	// Обычно используется библиотека для SQLite
	// sudo apt install build-essential
	// _ "github.com/mattn/go-sqlite3"

	// Если хотим без привязки к C библиотекам
	_ "modernc.org/sqlite"
)

// Todo struct, represent single task
type todo struct {
	id int
	task string
	owner string
	checked int
}



func main() {

	// Remove "todo.db" if file exists
	os.Remove("./todo.db")

	// Open database connection
	// github.com/mattn/go-sqlite3
	// db, err := sql.Open("sqlite3", "./todo.db")

	// modernc.org/sqlite
	db, err := sql.Open("sqlite", "./todo.db")

	// check connection result
	if err != nil {
		log.Fatal(err)
	}

	// clone connection
	defer db.Close()

	{ // Create table block
		// SQL statement to create table
		sqlStmt := `
		CREATE TABLE IF NOT EXISTS task (id INTEGER PRIMARY KEY AUTOINCREMENT, task TEXT, owner TEXT, checked INTEGER);
		`
		// Execute SQL statement
		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
		}
	}

	{	// Create records block
		// Begin transaction
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		// Prepare SQL statement than can be reused. Char "?" for SQLite, char "%" for MySQL, PostgreSQL
		stmt, err := tx.Prepare("INSERT INTO task(id, task, owner, checked) VALUES(?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		// close prepared statement before exiting program
		defer stmt.Close()

		// Create empty slice to store our todos
		tasks := []*todo{}
		// Create tasks
		tasks = append(tasks, &todo{id: 1, task: "Learn REST API", owner: "teacher", checked: 0})
		tasks = append(tasks, &todo{id: 2, task: "Make practice", owner: "students", checked: 0})

		for i := range tasks {
			// Insert records
			// Execute statement for each task
			_, err = stmt.Exec(tasks[i].id, tasks[i].task, tasks[i].owner, tasks[i].checked)
			if err != nil {
				log.Fatal(err)
			}
		}
		// Commit the transaction
		if err := tx.Commit(); err != nil {
			log.Fatal(err)
		}
	}

	{ // Read records block
		stmt, err := db.Prepare("SELECT id, task, owner FROM task WHERE checked = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		rows, err := stmt.Query(0)
		if err != nil {
			log.Fatal(err)
		}
		// Rows must be closed
		defer rows.Close()

		for rows.Next() {
			var id int
			var task string
			var owner string
			// use pointers to get data
			err = rows.Scan(&id, &task, &owner)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(id, task, owner)
		}
		// Check error, that can occur during iteration
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

	}

	{ 
		// Update block through UPDATE SQL statement 	
	}
	{ 
		// Delete block through DELETE SQL statement 	
	}

	
}