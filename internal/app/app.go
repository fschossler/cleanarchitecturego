package app

import (
	"log"
	"net/http"

	"github.com/fschossler/cleanarchitecturego/internal/controller"
	"github.com/gorilla/mux"
)

// App represents the application.
type App struct {
	Router *mux.Router
}

// NewApp initializes and configures the application.
func NewApp() *App {
	app := &App{
		Router: mux.NewRouter(),
	}

	// Initialize the database here if needed.
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize your dependencies and controllers.
	// Example:
	bookController := controller.NewBookController(app.Router, db)

	// Set up routes
	app.initRoutes()

	return app
}

// initRoutes sets up the application's routes.
func (a *App) initRoutes() {
	// Create your controller instances and initialize routes here.
	// Example:
	bookController.InitRoutes()
}

// StartServer starts the HTTP server.
func (a *App) StartServer(port string) {
	log.Printf("Server started on port %s\n", port)
	if err := http.ListenAndServe(port, a.Router); err != nil {
		log.Fatal(err)
	}
}
