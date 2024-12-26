package main

import (
	"log"
	"net/http"
	"os"

	svcs "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/services"
)

const (
	defaultPort              = "8081"
	defaultStorageURL        = "http://localhost:8083"
	defaultOrderProcessorURL = "http://localhost:8082"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	storageURL := os.Getenv("STORAGE_SERVICE_URL")
	if storageURL == "" {
		storageURL = defaultStorageURL
	}

	orderProcessorURL := os.Getenv("ORDER_PROCESSOR_URL")
	if orderProcessorURL == "" {
		orderProcessorURL = defaultOrderProcessorURL
	}

	service := svcs.NewBooksService(storageURL, orderProcessorURL)

	router := service.AddRoutes()

	log.Printf("Books Service running on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
