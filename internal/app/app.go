package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/fschossler/cleanarchitecturego/internal/controller"
	"github.com/fschossler/cleanarchitecturego/internal/repository"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App represents the application.
type App struct {
	Router     *mux.Router
	DB         *sql.DB
	Repository *repository.BookRepository
}

// NewApp initializes and configures the application.
func NewApp() *App {
	app := &App{
		Router: mux.NewRouter(),
	}

	// Initialize the PostgreSQL database connection
	db, err := sql.Open("postgres", "user=myuser dbname=mydatabase password=mypassword sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	app.DB = db

	// Initialize the book repository
	bookRepo := repository.NewBookRepository(db)
	app.Repository = bookRepo

	// Initialize the book controller with the app instance
	bookController := controller.NewBookController(app.Router, app.Repository) // Pass the router and repository

	// Set up routes for the book controller
	bookController.InitRoutes(app.Router)

	return app
}

// StartServer starts the HTTP server.
func (a *App) StartServer(port string) {
	log.Printf("Server started on port %s\n", port)
	if err := http.ListenAndServe(port, a.Router); err != nil {
		log.Fatal(err)
	}
}
