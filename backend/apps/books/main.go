package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	svcs "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/services"
)

const (
	defaultPort              = "8081"
	defaultStorageURL        = "http://localhost:8083"
	defaultOrderProcessorURL = "http://localhost:8082"
	defaultPageSize          = "20"
	defaultEnableCache       = "true"
	defaultLogLevel          = "info"
	defaultAPIVersion        = "v1"
)

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {
	port := getEnvWithDefault("PORT", defaultPort)
	storageURL := getEnvWithDefault("STORAGE_SERVICE_URL", defaultStorageURL)
	orderProcessorURL := getEnvWithDefault("ORDER_PROCESSOR_URL", defaultOrderProcessorURL)

	// Additional configuration from ConfigMap
	pageSize, err := strconv.Atoi(getEnvWithDefault("PAGE_SIZE", defaultPageSize))
	if err != nil {
		log.Fatalf("Invalid PAGE_SIZE: %v", err)
	}

	enableCache, err := strconv.ParseBool(getEnvWithDefault("ENABLE_CACHE", defaultEnableCache))
	if err != nil {
		log.Fatalf("Invalid ENABLE_CACHE: %v", err)
	}

	logLevel := getEnvWithDefault("LOG_LEVEL", defaultLogLevel)
	apiVersion := getEnvWithDefault("API_VERSION", defaultAPIVersion)

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	fmt.Printf("Starting Books Service with configuration:\n")
	fmt.Printf("- API Version: %s\n", apiVersion)
	fmt.Printf("- Page Size: %d\n", pageSize)
	fmt.Printf("- Cache Enabled: %v\n", enableCache)
	fmt.Printf("- Log Level: %s\n", logLevel)
	if dbUsername != "" {
		fmt.Printf("- DB Username: %s\n", dbUsername)
		if dbPassword != "" {
			fmt.Println("- DB Password: [REDACTED]")
		}
	} else {
		fmt.Println("- DB Credentials: not configured")
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
