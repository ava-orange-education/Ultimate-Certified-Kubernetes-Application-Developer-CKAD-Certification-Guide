package main

import (
	"log"
	"net/http"
	"os"

	opSvc "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/services"
)

const (
	defaultPort       = "8082"
	defaultStorageURL = "http://localhost:8083"
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

	ops := opSvc.NewOrderProcessingService(storageURL)

	router := ops.AddRoutes()

	log.Printf("Order Processor Service running on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
