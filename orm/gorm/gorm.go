package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	ID   uint `gorm:"primary_key"`
	Name string
	Age  uint
}

func main() {
	// create a database connection
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_ADDR"))
	panicOnErr(err)
	defer db.Close()

	// create table
	panicOnErr(db.AutoMigrate(&Person{}).Error)

	// insert a row
	panicOnErr(db.Create(&Person{Name: "L1212", Age: 1000}).Error)

	// query
	var person Person
	panicOnErr(db.First(&person, 1).Error) // find person with id 1
	log.Printf("1: %+v", person)

	panicOnErr(db.First(&person, "name = ?", "L1212").Error) // find person with name l1212
	log.Printf("name = L1212: %+v", person)

	// update
	panicOnErr(db.Model(&person).Update("Price", 2000).Error)
	log.Printf("update: %+v", person)

	// delete
	panicOnErr(db.Delete(&person).Error)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
