package main

import (
	"log"
	"net/http"

	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/repository"
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/services"
)

const StorageServicePort = "8083"

func main() {
	bookRepo := repository.NewBooksRepo()
	storageService := services.NewStorageService(bookRepo)

	router := storageService.AddRoutes()

	log.Printf("Storage Service running on :%s", StorageServicePort)
	if err := http.ListenAndServe(":"+StorageServicePort, router); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
