package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // must be imported
)

type Person struct {
	ID   int64
	Name string
	Age  int
}

func main() {

	// Start a connection
	db, err := sql.Open("mysql", os.Getenv("MYSQL_ADDR"))
	panicOnErr(err)
	defer db.Close()

	// Create a table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS person (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100) NULL,
		age INTEGER NULL
	)`)
	panicOnErr(err)

	// Insert
	wife := Person{Name: "Maria", Age: 18}
	row, err := db.Exec("INSERT INTO person (name, age) VALUES (?, ?)", wife.Name, wife.Age)
	panicOnErr(err)
	last, err := row.LastInsertId()
	panicOnErr(err)
	log.Printf("New row ID: %d", last)

	// Query
	rows, err := db.Query("SELECT * FROM person")
	panicOnErr(err)
	for rows.Next() {
		var newPerson Person
		rows.Scan(&newPerson.ID, &newPerson.Name, &newPerson.Age)
		log.Printf("Got person: %+v", newPerson)
	}
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
