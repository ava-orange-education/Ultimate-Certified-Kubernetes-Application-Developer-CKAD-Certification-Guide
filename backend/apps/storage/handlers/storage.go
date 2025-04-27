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

func (sh *StorageHandler) HealthLive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (sh *StorageHandler) HealthReady(w http.ResponseWriter, r *http.Request) {
	if !sh.canAccessStorage() {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Storage access failed"))
		return
	}

	if !sh.isCacheInitialized() {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Cache not initialized"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ready"))
}

func (sh *StorageHandler) HealthStartup(w http.ResponseWriter, r *http.Request) {
	if !sh.isInitialDataLoaded() {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Initial data loading in progress"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Started"))
}

func (sh *StorageHandler) canAccessStorage() bool {
	// Implementation to check storage access
	// For demonstration purposes, we'll return true
	return true
}

func (sh *StorageHandler) isCacheInitialized() bool {
	// Implementation to check if cache is initialized
	// For demonstration purposes, we'll return true
	return true
}

func (sh *StorageHandler) isInitialDataLoaded() bool {
	// Implementation to check if initial data is loaded
	// For demonstration purposes, we'll return true
	return true
}

func (sh *StorageHandler) AddRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Health check endpoints
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	router.Get("/health/live", sh.HealthLive)
	router.Get("/health/ready", sh.HealthReady)
	router.Get("/health/startup", sh.HealthStartup)

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
