package main

import (
	"context"
	"fmt"

	"github.com/amsen/go-rest-api-course/internal/database"
)

// Run - is going to be responsible for
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("starting up our application")

	dbConn, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database.")
		return err
	}

	if err := dbConn.Ping(context.Background()); err != nil {
		fmt.Println("Failed to ping the database.")
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
