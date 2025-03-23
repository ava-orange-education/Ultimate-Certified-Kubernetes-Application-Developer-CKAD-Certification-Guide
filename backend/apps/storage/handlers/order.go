package handlers

import (
	"encoding/json"
	"net/http"

	opModels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/models"
	httpPkg "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/pkg/http"
)

func (s *StorageHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	httpPkg.JSON(r.Context(), w, http.StatusOK, s.service.ListOrders())
}

func (s *StorageHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var order opModels.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	s.service.AddOrder(order)

	httpPkg.JSON(r.Context(), w, http.StatusCreated, order)
}

func (s *StorageHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("id")
	if orderID == "" {
		http.Error(w, "Order ID required", http.StatusBadRequest)
		return
	}

	order, exists := s.service.GetOrderByID(orderID)
	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	httpPkg.JSON(r.Context(), w, http.StatusOK, order)
}

func (s *StorageHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	var os opModels.UpdateOrderStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&os); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order, err := s.service.UpdateOrderStatus(os)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpPkg.JSON(r.Context(), w, http.StatusCreated, order)
}
