package repository

import (
	"database/sql"
	"log"

	"github.com/fschossler/cleanarchitecturego/internal/entity"
)

// BookRepository represents a repository for managing book data.
type BookRepository struct {
	DB *sql.DB
}

// NewBookRepository initializes and returns a new BookRepository.
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		DB: db,
	}
}

// CreateBook inserts a new book into the database.
func (r *BookRepository) CreateBook(book *entity.Book) (int, error) {
	query := "INSERT INTO books (title, author) VALUES ($1, $2) RETURNING id"
	var bookID int
	err := r.DB.QueryRow(query, book.Title, book.Author).Scan(&bookID)
	if err != nil {
		log.Printf("Failed to insert book: %v", err)
		return 0, err
	}
	return bookID, nil
}

// GetBookByID retrieves a book by its ID.
func (r *BookRepository) GetBookByID(id int) (*entity.Book, error) {
	query := "SELECT id, title, author FROM books WHERE id = $1"
	var book entity.Book
	err := r.DB.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		log.Printf("Failed to retrieve book: %v", err)
		return nil, err
	}
	return &book, nil
}

// UpdateBook updates an existing book in the database.
func (r *BookRepository) UpdateBook(book *entity.Book) error {
	query := "UPDATE books SET title = $2, author = $3 WHERE id = $1"
	_, err := r.DB.Exec(query, book.ID, book.Title, book.Author)
	if err != nil {
		log.Printf("Failed to update book: %v", err)
		return err
	}
	return nil
}

// DeleteBook deletes a book by its ID.
func (r *BookRepository) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete book: %v", err)
		return err
	}
	return nil
}

// GetBooks retrieves a list of all books from the database.
func (r *BookRepository) GetBooks() ([]entity.Book, error) {
	query := "SELECT id, title, author FROM books"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Printf("Failed to retrieve books: %v", err)
		return nil, err
	}
	defer rows.Close()

	var books []entity.Book
	for rows.Next() {
		var book entity.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			log.Printf("Failed to scan book: %v", err)
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Failed to retrieve books: %v", err)
		return nil, err
	}

	return books, nil
}
