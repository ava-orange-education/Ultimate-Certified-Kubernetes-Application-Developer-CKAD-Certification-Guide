package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/alcionai/clues/cluerr"
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
	var (
		or    opModels.CreateOrderRequest
		book  *models.Book
		order *opModels.Order

		status int
		err    error
	)

	if err := json.NewDecoder(r.Body).Decode(&or); err != nil {
		http.Error(w, "decoding create order request", http.StatusBadRequest)
		return
	}

	if book, status, err = oh.checkBookAvailability(r.Context(), or); err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	if order, status, err = oh.createOrder(r.Context(), book, or); err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	if status, err = oh.updateBookQuantity(r.Context(), book, or); err != nil {
		http.Error(w, err.Error(), status)
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

func (oh *OrdersHandler) checkBookAvailability(
	ctx context.Context,
	or opModels.CreateOrderRequest) (*models.Book, int, error) {
	resp, err := http.Get(fmt.Sprintf("%s/internal/books/get?id=%s", oh.storageServiceURL, or.BookID))
	if err != nil {
		return nil, http.StatusInternalServerError, cluerr.NewWC(ctx, "fetch book details")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, cluerr.NewWC(ctx, "fetch book details")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, cluerr.NewWC(ctx, "read get book response")
	}

	var book models.Book
	err = json.Unmarshal(b, &book)
	if err != nil {
		return nil, http.StatusInternalServerError, cluerr.NewWC(ctx, "unmarshall get book response")
	}

	if book.Quantity < 1 {
		return nil, http.StatusConflict, cluerr.NewWC(ctx, "book out of stock")
	}

	return &book, -1, nil
}

func (oh *OrdersHandler) createOrder(
	ctx context.Context,
	book *models.Book,
	or opModels.CreateOrderRequest) (*opModels.Order, int, error) {
	order := opModels.Order{
		ID:       uuid.NewString()[:7],
		BookID:   or.BookID,
		UserID:   or.UserID,
		Status:   "created",
		Price:    float64(or.Quantity) * book.Price,
		Quantity: or.Quantity,
		Created:  time.Now().Format(time.RFC3339),
	}

	createOrderBytes, err := json.Marshal(order)
	if err != nil {
		return nil, http.StatusInternalServerError, cluerr.NewWC(ctx, "marshall order")
	}

	req, err := http.NewRequest(http.MethodPost, oh.storageServiceURL+"/internal/orders/add", bytes.NewBuffer(createOrderBytes))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, http.StatusInternalServerError, cluerr.NewWC(ctx, "create add order request")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, resp.StatusCode, cluerr.NewWC(ctx, "add order")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, resp.StatusCode, cluerr.NewWC(ctx, "order not created")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, cluerr.NewWC(ctx, "read create order response")
	}

	var createdOrder opModels.Order
	err = json.Unmarshal(b, &createdOrder)
	if err != nil {
		return nil, http.StatusInternalServerError, cluerr.NewWC(ctx, "unmarshall create order response")
	}

	return &createdOrder, -1, nil
}

func (oh *OrdersHandler) updateBookQuantity(
	ctx context.Context,
	book *models.Book,
	or opModels.CreateOrderRequest) (int, error) {
	upr := opModels.UpdateBookQuantityRequest{
		BookID:   or.BookID,
		Quantity: book.Quantity - or.Quantity,
	}

	uprBytes, err := json.Marshal(upr)
	if err != nil {
		return http.StatusInternalServerError, cluerr.NewWC(ctx, "marshall update book quantity request")
	}

	req, err := http.NewRequest(http.MethodPut, oh.storageServiceURL+"/internal/books/update-quantity", bytes.NewBuffer(uprBytes))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return http.StatusInternalServerError, cluerr.NewWC(ctx, "create update book quantity request")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp.StatusCode, cluerr.NewWC(ctx, "update book quantity")
	}

	if resp.StatusCode != http.StatusCreated {
		return resp.StatusCode, cluerr.NewWC(ctx, "book quantity not updated")
	}

	return -1, nil
}
