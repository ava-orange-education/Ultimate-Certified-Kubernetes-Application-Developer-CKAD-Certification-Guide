package services

import (
	opHandlers "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/handlers"
	opRepo "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type OrderProcessingService struct {
	oph *opHandlers.OrdersHandler
}

func NewOrderProcessingService(opr *opRepo.OrderRepository) *OrderProcessingService {
	return &OrderProcessingService{
		oph: opHandlers.NewOrdersHandler(opr),
	}
}

func (ops *OrderProcessingService) AddRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/orders", func(r chi.Router) {
		r.Post("/create", ops.oph.CreateOrder)
		r.Get("/get", ops.oph.GetOrder)
		r.Put("/update-status", ops.oph.UpdateOrderStatus)
	})

	return router
}
