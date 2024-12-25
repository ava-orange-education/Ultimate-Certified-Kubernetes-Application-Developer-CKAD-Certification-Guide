package services

import (
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	StorageServiceURL  = "http://localhost:8083"
	OrderProcessingURL = "http://localhost:8082"
)

type BooksService struct {
	bh handlers.BooksHandler
}

func NewBooksService() *BooksService {
	return &BooksService{
		bh: *handlers.NewBooksHandler(StorageServiceURL, OrderProcessingURL),
	}
}

func (bs *BooksService) AddRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api/books", func(r chi.Router) {
		r.Get("/list", bs.bh.ListBooks)
		r.Post("/add", bs.bh.AddBook)
		r.Get("/details", bs.bh.GetBookDetails)
		r.Post("/purchase", bs.bh.InitiatePurchase)
	})

	return router
}
