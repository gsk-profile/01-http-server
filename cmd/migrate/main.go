package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	var option int

	fmt.Println("What do you want? 1 - Create DB, 2 - Delete DB")
	fmt.Print("Enter 1 or 2 : ")

	_, err := fmt.Scan(&option)
	if err != nil {
		fmt.Println("Error reading option", err)
		return
	}

	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@localhost:5432/mydb?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	switch option {
	case 1:
		fmt.Println("Creating DB")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case 2:
		fmt.Println("Deleting the DB")
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	default:
		fmt.Println("Invalid option")
	}

	log.Println("Migration complete")
}
