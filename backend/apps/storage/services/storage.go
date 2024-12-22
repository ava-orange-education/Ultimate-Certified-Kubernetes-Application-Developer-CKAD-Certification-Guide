package services

import (
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/handlers"
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/repository"

	"github.com/gorilla/mux"
)

type StorageService struct {
	sh *handlers.StorageHandler
}

func NewStorageService(br *repository.BooksRepo) *StorageService {
	return &StorageService{
		sh: handlers.NewStorageHandler(br),
	}
}

func (ss *StorageService) AddRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books/add", ss.sh.AddBook).Methods("POST")
	router.HandleFunc("/books/get", ss.sh.GetBook).Methods("GET")
	router.HandleFunc("/books/update", ss.sh.UpdateBook).Methods("PUT")
	router.HandleFunc("/internal/books/quantity", ss.sh.CheckQuantity).Methods("GET")
	router.HandleFunc("/internal/books/update-quantity", ss.sh.UpdateQuantity).Methods("PUT")

	return router
}
