package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/models"
)

// BatchProcessHandler handles batch processing of orders
func (h *OrdersHandler) BatchProcessHandler(w http.ResponseWriter, r *http.Request) {
	var req models.BatchProcessRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set default limit if not specified
	if req.Limit <= 0 {
		req.Limit = 100
	}

	// Simulate batch processing
	log.Printf("Processing batch of up to %d orders", req.Limit)

	// In a real implementation, this would process pending orders up to the limit
	response := models.BatchProcessResponse{
		ProcessedCount: req.Limit,
		Message:        "Batch processing completed successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
