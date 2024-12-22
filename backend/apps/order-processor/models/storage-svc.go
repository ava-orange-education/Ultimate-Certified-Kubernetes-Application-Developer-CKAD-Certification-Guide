package models

type StorageQuantityResponse struct {
	Quantity int `json:"quantity"`
}

type UpdateBookQuantityRequest struct {
	BookID   string `json:"book_id"`
	Quantity int    `json:"quantity"`
}
