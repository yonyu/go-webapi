package main

import (
	"context"
	"fmt"

	"github.com/yonyu/go-webapi/internal/database"
	"github.com/yonyu/go-webapi/internal/domain"
)

// Run - is going to be responsible for the instantiation
// and startup of the Go application
func Run() error {
	fmt.Println("Starting up the application")

	dbCurrent, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := dbCurrent.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate the database")
		return err
	}
	fmt.Println("Successfully connected and pinged the database")

	commentService := domain.NewService(dbCurrent)

	commentService.PostComment(
		context.Background(),
		domain.Comment{
			ID:     "095263fc-aacf-45e5-9b72-24926df586b9",
			Slug:   "Manual-test",
			Author: "John Smith",
			Body:   "Hello World",
		},
	)

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
