package services

import (
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/handlers"
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type StorageService struct {
	sh *handlers.StorageHandler
}

func NewStorageService(br *repository.BooksRepo) *StorageService {
	return &StorageService{
		sh: handlers.NewStorageHandler(br),
	}
}

func (ss *StorageService) AddRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/books", func(r chi.Router) {
		r.Get("/get", ss.sh.GetBook)
		r.Post("/add", ss.sh.AddBook)
		r.Put("/update", ss.sh.UpdateBook)
	})

	router.Route("/internal/books", func(r chi.Router) {
		r.Get("/quantity", ss.sh.CheckQuantity)
		r.Put("/update-quantity", ss.sh.UpdateQuantity)
	})

	return router
}
