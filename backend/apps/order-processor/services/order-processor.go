package services

import (
	"net/http"

	opHandlers "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (ops *OrderProcessingService) HealthLive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (ops *OrderProcessingService) HealthReady(w http.ResponseWriter, r *http.Request) {
	if !ops.canConnectToStorage() {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Storage connection failed"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ready"))
}

func (ops *OrderProcessingService) HealthStartup(w http.ResponseWriter, r *http.Request) {
	if !ops.isInitialSetupComplete() {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Initial setup in progress"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Started"))
}

func (ops *OrderProcessingService) canConnectToStorage() bool {
	// Implementation to check storage service connectivity
	// For demonstration purposes, we'll return true
	return true
}

func (ops *OrderProcessingService) isInitialSetupComplete() bool {
	// Implementation to check if initial setup is complete
	// For demonstration purposes, we'll return true
	return true
}

type OrderProcessingService struct {
	oph *opHandlers.OrdersHandler
}

func NewOrderProcessingService(storageURL string) *OrderProcessingService {
	return &OrderProcessingService{
		oph: opHandlers.NewOrdersHandler(storageURL),
	}
}

func (ops *OrderProcessingService) AddRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Health check endpoints
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	router.Get("/health/live", ops.HealthLive)
	router.Get("/health/ready", ops.HealthReady)
	router.Get("/health/startup", ops.HealthStartup)

	// API routes
	router.Route("/orders", func(r chi.Router) {
		r.Post("/create", ops.oph.CreateOrder)
		r.Put("/update-status", ops.oph.UpdateOrderStatus)

		// Batch processing endpoint for Jobs/CronJobs
		r.Post("/batch-process", ops.oph.BatchProcessHandler)
	})

	return router
}
