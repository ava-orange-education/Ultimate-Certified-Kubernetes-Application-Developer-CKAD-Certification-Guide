package handlers

import (
	"encoding/json"
	"net/http"

	booksmodels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/models"
	httpPkg "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/pkg/http"
)

func (s *StorageHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var book booksmodels.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	s.service.AddBook(book)

	httpPkg.JSON(r.Context(), w, http.StatusCreated, book)
}

func (s *StorageHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.URL.Query().Get("id")
	if bookID == "" {
		http.Error(w, "Book ID required", http.StatusBadRequest)
		return
	}

	book, exists := s.service.GetBookByID(bookID)
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	httpPkg.JSON(r.Context(), w, http.StatusOK, book)
}

func (s *StorageHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	httpPkg.JSON(r.Context(), w, http.StatusOK, s.service.GetBooks())
}

func (s *StorageHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book booksmodels.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := s.service.UpdateBook(book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpPkg.JSON(r.Context(), w, http.StatusCreated, book)
}
