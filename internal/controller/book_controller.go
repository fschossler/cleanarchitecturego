package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fschossler/cleanarchitecturego/internal/entity"
	"github.com/fschossler/cleanarchitecturego/internal/repository"
	"github.com/gorilla/mux"
)

// BookController represents the controller for book-related operations.
type BookController struct {
	Router     *mux.Router
	Repository *repository.BookRepository
}

// NewBookController initializes and returns a new BookController.
func NewBookController(router *mux.Router, repo *repository.BookRepository) *BookController {
	return &BookController{
		Router:     router,
		Repository: repo,
	}
}

func (c *BookController) InitRoutes(router *mux.Router) {
	c.Router = router // Set the router
	c.Router.HandleFunc("/api/books", c.CreateBook).Methods("POST")
	c.Router.HandleFunc("/api/books", c.GetBooks).Methods("GET")
	c.Router.HandleFunc("/api/books/{id}", c.GetBookByID).Methods("GET")
	c.Router.HandleFunc("/api/books/{id}", c.UpdateBook).Methods("PUT")
	c.Router.HandleFunc("/api/books/{id}", c.DeleteBook).Methods("DELETE")
}

// CreateBook handles the creation of a new book.
func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookID, err := c.Repository.CreateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the newly created book's ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": bookID})
}

// GetBooks retrieves a list of all books.
func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.Repository.GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

// GetBookByID retrieves a book by its ID.
func (c *BookController) GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := c.Repository.GetBookByID(bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook handles the updating of a book.
func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book.ID = bookID
	if err := c.Repository.UpdateBook(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteBook deletes a book by its ID.
func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.Repository.DeleteBook(bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
