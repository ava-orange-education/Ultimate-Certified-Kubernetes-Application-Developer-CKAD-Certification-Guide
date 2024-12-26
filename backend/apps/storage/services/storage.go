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

func NewStorageService(
	br *repository.BooksRepo,
	or *repository.OrderRepository) *StorageService {
	return &StorageService{
		sh: handlers.NewStorageHandler(br, or),
	}
}

func (ss *StorageService) AddRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/internal/books", func(r chi.Router) {
		r.Get("/list", ss.sh.ListBooks)
		r.Get("/get", ss.sh.GetBook)
		r.Post("/add", ss.sh.AddBook)
		r.Put("/update", ss.sh.UpdateBook)
		r.Get("/quantity", ss.sh.CheckQuantity)
		r.Put("/update-quantity", ss.sh.UpdateQuantity)
	})

	router.Route("/internal/orders", func(r chi.Router) {
		r.Get("/list", ss.sh.ListOrders)
		r.Get("/get", ss.sh.GetOrder)
		r.Post("/add", ss.sh.AddOrder)
		r.Put("/update-status", ss.sh.UpdateOrderStatus)
	})

	return router
}
