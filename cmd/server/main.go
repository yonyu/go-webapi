package main

import (
	"fmt"

	"github.com/yonyu/go-webapi/internal/database"
	"github.com/yonyu/go-webapi/internal/domain"
	transportHttp "github.com/yonyu/go-webapi/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(commentService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go Rest Web API")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
