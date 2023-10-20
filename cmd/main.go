package main

import "github.com/fschossler/cleanarchitecturego/internal/app"

func main() {
	// Initialize the application
	myApp := app.NewApp()

	// Start the HTTP server
	port := ":8080"
	myApp.StartServer(port)
}
