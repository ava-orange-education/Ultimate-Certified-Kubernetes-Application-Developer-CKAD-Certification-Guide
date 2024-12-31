package services

import (
	opHandlers "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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

	// API routes
	router.Route("/orders", func(r chi.Router) {
		r.Post("/create", ops.oph.CreateOrder)
		r.Put("/update-status", ops.oph.UpdateOrderStatus)

		// Batch processing endpoint for Jobs/CronJobs
		r.Post("/batch-process", ops.oph.BatchProcessHandler)
	})

	return router
}
