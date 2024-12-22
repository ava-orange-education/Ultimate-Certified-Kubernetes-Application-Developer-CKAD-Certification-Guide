package services

import (
	opHandlers "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/handlers"
	opRepo "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/repository"

	"github.com/gorilla/mux"
)

type OrderProcessingService struct {
	oph *opHandlers.OrdersHandler
}

func NewOrderProcessingService(opr *opRepo.OrderRepository) *OrderProcessingService {
	return &OrderProcessingService{
		oph: opHandlers.NewOrdersHandler(opr),
	}
}

func (ops *OrderProcessingService) AddRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/orders/create", ops.oph.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/get", ops.oph.GetOrder).Methods("GET")
	router.HandleFunc("/orders/update-status", ops.oph.UpdateOrderStatus).Methods("PUT")

	return router
}
