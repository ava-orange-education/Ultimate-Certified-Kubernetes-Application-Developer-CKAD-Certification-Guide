package models

type BatchProcessRequest struct {
	Limit int `json:"limit"`
}

type BatchProcessResponse struct {
	ProcessedCount int    `json:"processed_count"`
	Message        string `json:"message"`
}
