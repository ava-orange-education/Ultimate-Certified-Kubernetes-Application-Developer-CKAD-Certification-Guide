package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BooksHandler struct {
	storageServiceURL  string
	orderProcessingURL string
}

func NewBooksHandler(storageServiceURL, orderProcessingURL string) *BooksHandler {
	return &BooksHandler{
		storageServiceURL:  storageServiceURL,
		orderProcessingURL: orderProcessingURL,
	}
}

func (bh *BooksHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp, err := http.Get(bh.storageServiceURL + "/books/list")
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	io.Copy(w, resp.Body)
}

func (bh *BooksHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Forward request to storage service
	resp, err := http.Post(bh.storageServiceURL+"/books/add", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Failed to add book", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func (bh *BooksHandler) InitiatePurchase(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var purchaseRequest struct {
		BookID string `json:"book_id"`
		UserID string `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&purchaseRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check book availability
	resp, err := http.Get(fmt.Sprintf("%s/internal/books/quantity?id=%s", bh.storageServiceURL, purchaseRequest.BookID))
	if err != nil {
		http.Error(w, "Failed to check book availability", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Book not available", resp.StatusCode)
		return
	}

	// Forward purchase request to order processing
	orderData, err := json.Marshal(purchaseRequest)
	if err != nil {
		http.Error(w, "Failed to process purchase", http.StatusInternalServerError)
		return
	}

	resp, err = http.Post(bh.orderProcessingURL+"/orders/create", "application/json", bytes.NewBuffer(orderData))
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func (bh *BooksHandler) GetBookDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bookID := r.URL.Query().Get("id")
	if bookID == "" {
		http.Error(w, "Book ID required", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(fmt.Sprintf("%s/books/get?id=%s", bh.storageServiceURL, bookID))
	if err != nil {
		http.Error(w, "Failed to fetch book details", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
