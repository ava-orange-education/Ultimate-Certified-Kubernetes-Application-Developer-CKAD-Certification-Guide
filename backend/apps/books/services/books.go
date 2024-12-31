package services

import (
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type BooksService struct {
	bh handlers.BooksHandler
}

func NewBooksService(storageURL, orderProcessorURL string) *BooksService {
	return &BooksService{
		bh: *handlers.NewBooksHandler(storageURL, orderProcessorURL),
	}
}

func (bs *BooksService) AddRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // In production, specify exact origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// API routes
	router.Route("/api/books", func(r chi.Router) {
		r.Get("/list", bs.bh.ListBooks)
		r.Post("/add", bs.bh.AddBook)
		r.Get("/details", bs.bh.GetBookDetails)
		r.Post("/purchase", bs.bh.InitiatePurchase)
	})

	return router
}
