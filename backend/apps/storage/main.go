package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/handlers"
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/repository"
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/services"
)

const (
	defaultPort = "8083"
	logDir      = "/app/logs"
	logFile     = "storage-service.log"
	cacheDir    = "/app/cache"
	dataDir     = "/data"
)

func setupLogging() (*os.File, error) {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %v", err)
	}

	logPath := filepath.Join(logDir, logFile)
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf("Logging initialized to %s", logPath)
	return f, nil
}

func setupCache() error {
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %v", err)
	}
	log.Printf("Cache directory initialized at %s", cacheDir)
	return nil
}

func setupDataDir() error {
	dir := os.Getenv("DATA_DIR")
	if dir == "" {
		dir = dataDir
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %v", err)
	}
	log.Printf("Data directory initialized at %s", dir)
	return nil
}

func main() {
	logFile, err := setupLogging()
	if err != nil {
		log.Printf("Warning: Could not set up file logging: %v", err)
	} else {
		defer logFile.Close()
	}

	if err := setupCache(); err != nil {
		log.Printf("Warning: Could not set up cache directory: %v", err)
	}

	if err := setupDataDir(); err != nil {
		log.Printf("Warning: Could not set up data directory: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	bookRepo := repository.NewBooksRepo()
	orderRepo := repository.NewOrderRepository()

	storageService := services.NewStorageService(bookRepo, orderRepo)
	storageHandler := handlers.NewStorageHandler(storageService)

	cachedBooksHandler := handlers.NewCachedBooksHandler(storageService)

	router := storageHandler.AddRoutes()
	cachedBooksHandler.AddCachedRoutes(router)

	log.Printf("Storage Service running on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
