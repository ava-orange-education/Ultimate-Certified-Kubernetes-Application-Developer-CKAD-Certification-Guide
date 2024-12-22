package handlers

import (
	"encoding/json"
	"net/http"

	storagemodels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/models"
)

func (s *StorageHandler) CheckQuantity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bookID := r.URL.Query().Get("id")
	if bookID == "" {
		http.Error(w, "Book ID required", http.StatusBadRequest)
		return
	}

	book, exists := s.br.GetBookByID(bookID)
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"quantity": book.Quantity})
}

func (s *StorageHandler) UpdateQuantity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updateReq storagemodels.UpdateBookQuantityRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	book, err := s.br.UpdateBookQuantity(updateReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(book)
}
