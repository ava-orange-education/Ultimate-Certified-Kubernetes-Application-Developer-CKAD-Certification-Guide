package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/models"
	opModels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/models"
	httpPkg "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/pkg/http"

	"github.com/google/uuid"
)

const (
	BookSellerID = "test-seller"
	BookBuyerID  = "test-buyer"
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
	resp, err := http.Get(bh.storageServiceURL + "/internal/books/list")
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read storage book list response", http.StatusInternalServerError)
		return
	}

	var books []models.Book
	err = json.Unmarshal(b, &books)
	if err != nil {
		http.Error(w, "Failed to unmarshall book list response", http.StatusInternalServerError)
		return
	}

	httpPkg.JSON(r.Context(), w, resp.StatusCode, books)
}

func (bh *BooksHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read add book request", http.StatusInternalServerError)
		return
	}

	var bookReq models.Book
	err = json.Unmarshal(b, &bookReq)
	if err != nil {
		http.Error(w, "Failed to unmarshall add book request", http.StatusInternalServerError)
		return
	}

	bookReq.ID = uuid.NewString()[:7]
	bookReq.SellerID = BookSellerID

	b, err = json.Marshal(bookReq)
	if err != nil {
		http.Error(w, "Failed to marshall add book request", http.StatusInternalServerError)
		return
	}

	// Forward request to storage service
	resp, err := http.Post(bh.storageServiceURL+"/internal/books/add", "application/json", bytes.NewBuffer(b))
	if err != nil {
		http.Error(w, "Failed to add book", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read storage add book response", http.StatusInternalServerError)
		return
	}

	var book models.Book
	err = json.Unmarshal(b, &book)
	if err != nil {
		http.Error(w, "Failed to unmarshall add book response", http.StatusInternalServerError)
		return
	}

	httpPkg.JSON(r.Context(), w, resp.StatusCode, book)
}

func (bh *BooksHandler) InitiatePurchase(w http.ResponseWriter, r *http.Request) {
	var pReq opModels.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&pReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	pReq.UserID = BookBuyerID

	// Forward purchase request to order processing
	orderData, err := json.Marshal(pReq)
	if err != nil {
		http.Error(w, "Failed to process purchase", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(bh.orderProcessingURL+"/orders/create", "application/json", bytes.NewBuffer(orderData))
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read storage create order response", http.StatusInternalServerError)
		return
	}

	var order opModels.Order
	err = json.Unmarshal(b, &order)
	if err != nil {
		http.Error(w, "Failed to unmarshall create order response", http.StatusInternalServerError)
		return
	}

	httpPkg.JSON(r.Context(), w, resp.StatusCode, order)
}

func (bh *BooksHandler) GetBookDetails(w http.ResponseWriter, r *http.Request) {
	bookID := r.URL.Query().Get("id")
	if bookID == "" {
		http.Error(w, "Book ID required", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(fmt.Sprintf("%s/internal/books/get?id=%s", bh.storageServiceURL, bookID))
	if err != nil {
		http.Error(w, "Failed to fetch book details", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to fetch book details", resp.StatusCode)
		return
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read storage get book response", http.StatusInternalServerError)
		return
	}

	var book models.Book
	err = json.Unmarshal(b, &book)
	if err != nil {
		http.Error(w, "Failed to unmarshall get book response", http.StatusInternalServerError)
		return
	}

	httpPkg.JSON(r.Context(), w, resp.StatusCode, book)
}

func HealthLive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func HealthReady(w http.ResponseWriter, r *http.Request) {
	if !canConnectToDatabase() {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Database connection failed"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ready"))
}

func canConnectToDatabase() bool {
	// Implementation to check database connectivity
	// For demonstration purposes, we'll return true
	return true
}
