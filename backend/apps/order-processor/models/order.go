package models

type Order struct {
	ID      string  `json:"id"`
	BookID  string  `json:"book_id"`
	UserID  string  `json:"user_id"`
	Status  string  `json:"status"`
	Price   float64 `json:"price"`
	Created string  `json:"created"`
}

type CreateOrderRequest struct {
	BookID string `json:"book_id"`
	UserID string `json:"user_id"`
}

type UpdateOrderStatusRequest struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}
