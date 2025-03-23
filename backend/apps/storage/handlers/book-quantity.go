package handlers

import (
	"encoding/json"
	"net/http"

	storagemodels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/models"
	httpPkg "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/pkg/http"
)

func (s *StorageHandler) CheckQuantity(w http.ResponseWriter, r *http.Request) {
	bookID := r.URL.Query().Get("id")
	if bookID == "" {
		http.Error(w, "Book ID required", http.StatusBadRequest)
		return
	}

	quantity, exists := s.service.CheckQuantity(bookID)
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	httpPkg.JSON(r.Context(), w, http.StatusOK, map[string]int{"quantity": quantity})
}

func (s *StorageHandler) UpdateQuantity(w http.ResponseWriter, r *http.Request) {
	var updateReq storagemodels.UpdateBookQuantityRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	book, err := s.service.UpdateBookQuantity(updateReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpPkg.JSON(r.Context(), w, http.StatusCreated, book)
}
