package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Order Processor Service running on :%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-stop

	log.Println("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Forced shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}
