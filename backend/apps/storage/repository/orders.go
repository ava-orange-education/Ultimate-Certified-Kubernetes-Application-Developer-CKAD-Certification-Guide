package repository

import (
	"errors"
	"log"
	"sync"

	opModels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/models"
)

type OrderRepository struct {
	orders map[string]opModels.Order
	mu     sync.RWMutex
	pm     *PersistenceManager
}

func NewOrderRepository() *OrderRepository {
	pm := NewPersistenceManager()
	orders, err := pm.LoadOrders()
	if err != nil {
		log.Printf("Warning: Could not load orders from disk: %v", err)
		orders = make(map[string]opModels.Order)
	}

	return &OrderRepository{
		orders: orders,
		pm:     pm,
	}
}

func (or *OrderRepository) AddOrder(order opModels.Order) {
	or.mu.Lock()
	or.orders[order.ID] = order
	or.mu.Unlock()

	if err := or.pm.SaveOrders(or.orders); err != nil {
		log.Printf("Warning: Could not save orders to disk: %v", err)
	}
}

func (or *OrderRepository) GetOrderByID(orderID string) (opModels.Order, bool) {
	or.mu.Lock()
	order, exists := or.orders[orderID]
	or.mu.Unlock()

	return order, exists
}

func (or *OrderRepository) ListOrders() []opModels.Order {
	or.mu.Lock()
	defer or.mu.Unlock()

	orders := make([]opModels.Order, 0)
	for _, o := range or.orders {
		orders = append(orders, o)
	}

	return orders
}

func (or *OrderRepository) UpdateOrderStatus(uosr opModels.UpdateOrderStatusRequest) (opModels.Order, error) {
	if _, exists := or.GetOrderByID(uosr.OrderID); !exists {
		return opModels.Order{}, errors.New("order not found")
	}

	or.mu.Lock()
	order := or.orders[uosr.OrderID]
	order.Status = uosr.Status
	or.orders[uosr.OrderID] = order
	or.mu.Unlock()

	if err := or.pm.SaveOrders(or.orders); err != nil {
		log.Printf("Warning: Could not save orders to disk: %v", err)
	}

	return order, nil
}
