package repository

import (
	"errors"
	"sync"

	opModels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/models"
)

type OrderRepository struct {
	orders map[string]opModels.Order
	mu     sync.RWMutex
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]opModels.Order),
	}
}

func (or *OrderRepository) AddOrder(order opModels.Order) {
	or.mu.Lock()
	or.orders[order.ID] = order
	or.mu.Unlock()
}

func (or *OrderRepository) GetOrder(orderID string) (opModels.Order, bool) {
	or.mu.Lock()
	order, exists := or.orders[orderID]
	or.mu.Unlock()

	return order, exists
}

func (or *OrderRepository) UpdateOrderStatus(uosr opModels.UpdateOrderStatusRequest) (opModels.Order, error) {
	if _, exists := or.GetOrder(uosr.OrderID); !exists {
		return opModels.Order{}, errors.New("order not found")
	}

	or.mu.Lock()
	order := or.orders[uosr.OrderID]
	order.Status = uosr.Status
	or.orders[uosr.OrderID] = order
	or.mu.Unlock()

	return order, nil
}
