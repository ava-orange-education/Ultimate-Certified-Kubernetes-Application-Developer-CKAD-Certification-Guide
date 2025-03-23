package handlers

import (
	"net/http"

	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type StorageHandler struct {
	service *services.StorageService
}

func NewStorageHandler(service *services.StorageService) *StorageHandler {
	return &StorageHandler{service: service}
}

func (sh *StorageHandler) AddRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	router.Route("/internal/books", func(r chi.Router) {
		r.Get("/list", sh.ListBooks)
		r.Get("/get", sh.GetBook)
		r.Post("/add", sh.AddBook)
		r.Put("/update", sh.UpdateBook)
		r.Get("/quantity", sh.CheckQuantity)
		r.Put("/update-quantity", sh.UpdateQuantity)
	})

	router.Route("/internal/orders", func(r chi.Router) {
		r.Get("/list", sh.ListOrders)
		r.Get("/get", sh.GetOrder)
		r.Post("/add", sh.AddOrder)
		r.Put("/update-status", sh.UpdateOrderStatus)
	})

	return router
}
