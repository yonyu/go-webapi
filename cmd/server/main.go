package main

import (
	"context"
	"fmt"

	"github.com/yonyu/go-webapi/internal/db"
	"github.com/yonyu/go-webapi/internal/domain"
)

// Run - is going to be responsible for the instantiation
// and startup of the Go application
func Run() error {
	fmt.Println("Starting up the application")

	database, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := database.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate the database")
		return err
	}
	fmt.Println("Successfully connected and pinged the database")

	commentService := domain.NewService(database)
	fmt.Println(commentService.GetComment(
		context.Background(),
		"84f279ee-5aef-4ddb-8ae7-0d561a7944b2",
	))

	return nil
}

func main() {
	fmt.Println("Go Rest Web API")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
