package repository

import (
	"errors"
	"sync"

	booksmodels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/models"
	storagemodels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/models"
)

type BooksRepo struct {
	books map[string]booksmodels.Book
	mu    sync.RWMutex
}

func NewBooksRepo() *BooksRepo {
	return &BooksRepo{
		books: make(map[string]booksmodels.Book),
	}
}

func (br *BooksRepo) AddBook(book booksmodels.Book) {
	br.mu.Lock()
	br.books[book.ID] = book
	br.mu.Unlock()
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
