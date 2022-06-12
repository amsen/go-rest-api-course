package main

import (
	"fmt"
	"time"

	"github.com/amsen/go-rest-api-course/internal/database"
)

// Run - is going to be responsible for
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("starting up our application")

	// Database starting up before the application fails the migrations
	// need to solve this in docker compose. this is a temporary fix
	time.Sleep(10 * time.Second)

	dbConn, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database.")
		return err
	}

	if err := dbConn.MigrateDB(); err != nil {
		fmt.Print("Failed to run migrations on the database")
		return err
	}

	fmt.Println("Sucessfully pinged the database..")

	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
