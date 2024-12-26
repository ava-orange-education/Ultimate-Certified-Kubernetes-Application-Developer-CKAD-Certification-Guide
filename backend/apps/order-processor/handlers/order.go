package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/models"
	opModels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/models"
	httpPkg "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/pkg/http"

	"github.com/google/uuid"
)

const StorageServiceURL = "http://localhost:8083"

type OrdersHandler struct {
	storageServiceURL string
}

func NewOrdersHandler() *OrdersHandler {
	return &OrdersHandler{
		storageServiceURL: StorageServiceURL,
	}
}

func (oh *OrdersHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var or opModels.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&or); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(fmt.Sprintf("%s/internal/books/get?id=%s", oh.storageServiceURL, or.BookID))
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to get book details: %v", err),
			http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read storage add book response", http.StatusInternalServerError)
		return
	}

	var book models.Book
	err = json.Unmarshal(b, &book)
	if err != nil {
		http.Error(w, "Failed to unmarshall add book response", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Book not available", resp.StatusCode)
		return
	}

	if book.Quantity < 1 {
		http.Error(w, "Book out of stock", http.StatusConflict)
		return
	}

	order := opModels.Order{
		ID:      uuid.NewString()[:7],
		BookID:  or.BookID,
		UserID:  or.UserID,
		Status:  "created",
		Price:   float64(or.Quantity) * book.Price,
		Created: time.Now().Format(time.RFC3339),
	}

	createOrderBytes, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "Failed to marshall order", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(http.MethodPost, oh.storageServiceURL+"/internal/orders/add", bytes.NewBuffer(createOrderBytes))
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

	if resp.StatusCode != http.StatusCreated {
		http.Error(w, "Order not created", resp.StatusCode)
		return
	}

	upr := opModels.UpdateBookQuantityRequest{
		BookID:   or.BookID,
		Quantity: book.Quantity - or.Quantity,
	}

	updateBody, err := json.Marshal(upr)
	if err != nil {
		http.Error(w, "Failed to marshall update book quantity request", http.StatusInternalServerError)
		return
	}

	req, err = http.NewRequest(http.MethodPut, oh.storageServiceURL+"/internal/books/update-quantity", bytes.NewBuffer(updateBody))
	if err != nil {
		http.Error(w, "Failed to create update request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		http.Error(w, "Failed to update book quantity", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusCreated {
		http.Error(w, "Book qunatity not updated", resp.StatusCode)
		return
	}

	httpPkg.JSON(r.Context(), w, http.StatusCreated, order)
}

func (oh *OrdersHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	var uosr opModels.UpdateOrderStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&uosr); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	uosrBytes, err := json.Marshal(uosr)
	if err != nil {
		http.Error(w, "Failed to marshall update order status request", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(http.MethodPut, oh.storageServiceURL+"/internal/orders/update-status", bytes.NewBuffer(uosrBytes))
	if err != nil {
		http.Error(w, "Failed to update order status", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to update order status", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		http.Error(w, "Order status not updated", resp.StatusCode)
		return
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read update order status response", resp.StatusCode)
		return
	}

	var order opModels.Order
	err = json.Unmarshal(b, &order)
	if err != nil {
		http.Error(w, "Failed to unmarshall update order status response", resp.StatusCode)
		return
	}

	httpPkg.JSON(r.Context(), w, http.StatusCreated, order)
}
