package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

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

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Storage Service running on :%s", port)
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
