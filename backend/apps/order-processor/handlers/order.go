package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	opModels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/models"
	opRepo "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/repository"
)

type OrdersHandler struct {
	or                *opRepo.OrderRepository
	storageServiceURL string
}

func NewOrdersHandler(or *opRepo.OrderRepository) *OrdersHandler {
	return &OrdersHandler{
		or: or,
	}
}

func (oh *OrdersHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var or opModels.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&or); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(fmt.Sprintf("%s/internal/books/quantity?id=%s", oh.storageServiceURL, or.BookID))
	if err != nil {
		http.Error(w, "Failed to verify book quantity", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Book not available", resp.StatusCode)
		return
	}

	var qr opModels.StorageQuantityResponse
	if err := json.NewDecoder(resp.Body).Decode(&qr); err != nil {
		http.Error(w, "Failed to parse quantity response", http.StatusInternalServerError)
		return
	}

	if qr.Quantity < 1 {
		http.Error(w, "Book out of stock", http.StatusConflict)
		return
	}

	upr := opModels.UpdateBookQuantityRequest{
		BookID:   or.BookID,
		Quantity: qr.Quantity - 1,
	}

	updateBody, err := json.Marshal(upr)
	if err != nil {
		http.Error(w, "Failed to process order", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(http.MethodPut, oh.storageServiceURL+"/internal/books/update-quantity", bytes.NewBuffer(updateBody))
	if err != nil {
		http.Error(w, "Failed to create update request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		http.Error(w, "Failed to update book quantity", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	order := opModels.Order{
		ID:      "",
		BookID:  or.BookID,
		UserID:  or.UserID,
		Status:  "created",
		Created: time.Now().Format(time.RFC3339),
	}

	oh.or.AddOrder(order)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (oh *OrdersHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	orderID := r.URL.Query().Get("id")
	if orderID == "" {
		http.Error(w, "Order ID required", http.StatusBadRequest)
		return
	}

	order, exists := oh.or.GetOrder(orderID)
	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func (oh *OrdersHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var uosr opModels.UpdateOrderStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&uosr); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order, err := oh.or.UpdateOrderStatus(uosr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(order)
}
