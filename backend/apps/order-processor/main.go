package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	opSvc "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/services"
)

const (
	defaultPort       = "8082"
	defaultStorageURL = "http://localhost:8083"
	defaultOrdersDir  = "/orders"
)

func setupOrdersDir() error {
	dir := os.Getenv("ORDER_DATA_DIR")
	if dir == "" {
		dir = defaultOrdersDir
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create orders directory: %v", err)
	}
	log.Printf("Orders directory initialized at %s", dir)
	return nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	storageURL := os.Getenv("STORAGE_SERVICE_URL")
	if storageURL == "" {
		storageURL = defaultStorageURL
	}

	if err := setupOrdersDir(); err != nil {
		log.Printf("Warning: Could not set up orders directory: %v", err)
	}

	ops := opSvc.NewOrderProcessingService(storageURL)

	router := ops.AddRoutes()

	log.Printf("Order Processor Service running on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
