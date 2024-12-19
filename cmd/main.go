package main

import "fmt"

// Run - is going to be responsible for the instantiation
// and startup of the Go application
func Run() error {
	fmt.Println("Starting up the application")
	return nil
}

func main() {
	fmt.Println("Go Rest Web API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
