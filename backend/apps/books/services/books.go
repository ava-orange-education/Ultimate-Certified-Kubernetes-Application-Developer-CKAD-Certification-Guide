package services

import (
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/handlers"

	"github.com/gorilla/mux"
)

const (
	StorageServiceURL  = "http://storage-service:8083"
	OrderProcessingURL = "http://order-processing:8082"
)

type BooksService struct {
	bh handlers.BooksHandler
}

func NewBooksService() *BooksService {
	return &BooksService{
		bh: *handlers.NewBooksHandler(StorageServiceURL, OrderProcessingURL),
	}
}

func (bs *BooksService) AddRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/books/list", bs.bh.ListBooks).Methods("GET")
	router.HandleFunc("/api/books/add", bs.bh.AddBook).Methods("POST")
	router.HandleFunc("/api/books/details", bs.bh.GetBookDetails).Methods("GET")
	router.HandleFunc("/api/purchase", bs.bh.InitiatePurchase).Methods("POST")

	return router
}
