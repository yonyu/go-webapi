package main

import (
	"context"
	"fmt"

	"github.com/yonyu/go-webapi/internal/db"
)

// Run - is going to be responsible for the instantiation
// and startup of the Go application
func Run() error {
	fmt.Println("Starting up the application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("Successfully connected and pinged the database")

	return nil
}

func main() {
	fmt.Println("Go Rest Web API")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
