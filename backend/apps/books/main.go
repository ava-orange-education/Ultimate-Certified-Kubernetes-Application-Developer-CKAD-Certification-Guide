package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Books Service running on :%s", port)
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
