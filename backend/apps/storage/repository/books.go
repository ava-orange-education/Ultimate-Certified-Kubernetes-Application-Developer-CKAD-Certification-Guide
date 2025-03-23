package repository

import (
	"errors"
	"log"
	"sync"

	booksmodels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/models"
	storagemodels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/models"
)

type BooksRepo struct {
	books map[string]booksmodels.Book
	mu    sync.RWMutex
	pm    *PersistenceManager
}

func NewBooksRepo() *BooksRepo {
	pm := NewPersistenceManager()
	books, err := pm.LoadBooks()
	if err != nil {
		log.Printf("Warning: Could not load books from disk: %v", err)
		books = make(map[string]booksmodels.Book)
	}

	return &BooksRepo{
		books: books,
		pm:    pm,
	}
}

func (br *BooksRepo) AddBook(book booksmodels.Book) {
	br.mu.Lock()
	br.books[book.ID] = book
	br.mu.Unlock()

	if err := br.pm.SaveBooks(br.books); err != nil {
		log.Printf("Warning: Could not save books to disk: %v", err)
	}
}

func (br *BooksRepo) UpdateBook(book booksmodels.Book) error {
	if _, exists := br.GetBookByID(book.ID); !exists {
		return errors.New("book not found")
	}

	br.AddBook(book)

	return nil
}

func (br *BooksRepo) UpdateBookQuantity(req storagemodels.UpdateBookQuantityRequest) (booksmodels.Book, error) {
	if _, exists := br.GetBookByID(req.BookID); !exists {
		return booksmodels.Book{}, errors.New("book not found")
	}

	br.mu.Lock()
	book := br.books[req.BookID]
	book.Quantity = req.Quantity
	br.books[req.BookID] = book
	br.mu.Unlock()

	if err := br.pm.SaveBooks(br.books); err != nil {
		log.Printf("Warning: Could not save books to disk: %v", err)
	}

	return book, nil
}

func (br *BooksRepo) GetBookByID(id string) (booksmodels.Book, bool) {
	br.mu.Lock()
	book, exists := br.books[id]
	br.mu.Unlock()

	return book, exists
}

func (br *BooksRepo) GetBooks() []booksmodels.Book {
	br.mu.Lock()
	defer br.mu.Unlock()

	bookList := make([]booksmodels.Book, 0)

	for _, b := range br.books {
		bookList = append(bookList, b)
	}

	return bookList
}
