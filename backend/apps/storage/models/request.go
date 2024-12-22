package models

type UpdateBookQuantityRequest struct {
	BookID   string `json:"book_id"`
	Quantity int    `json:"quantity"`
}
