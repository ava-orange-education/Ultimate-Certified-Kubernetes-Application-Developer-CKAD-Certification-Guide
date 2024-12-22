package models

// Book represents a book in the marketplace
type Book struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	SellerID    string  `json:"seller_id"`
	Description string  `json:"description,omitempty"`
}
