package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/repository"
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/services"
)

const defaultPort = "8083"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	bookRepo := repository.NewBooksRepo()
	orderRepo := repository.NewOrderRepository()
	storageService := services.NewStorageService(bookRepo, orderRepo)

	router := storageService.AddRoutes()

	log.Printf("Storage Service running on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
