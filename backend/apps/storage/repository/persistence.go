package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	booksmodels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/models"
	opModels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/models"
)

const (
	defaultDataDir = "/data"
	booksFile      = "books.json"
	ordersFile     = "orders.json"
)

type PersistenceManager struct {
	dataDir string
	mu      sync.Mutex
}

func NewPersistenceManager() *PersistenceManager {
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = defaultDataDir
	}

	return &PersistenceManager{
		dataDir: dataDir,
	}
}

func (pm *PersistenceManager) SaveBooks(books map[string]booksmodels.Book) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	data, err := json.Marshal(books)
	if err != nil {
		return fmt.Errorf("error marshaling books: %v", err)
	}

	filePath := filepath.Join(pm.dataDir, booksFile)
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("error writing books file: %v", err)
	}

	log.Printf("Books saved to %s", filePath)
	return nil
}

func (pm *PersistenceManager) LoadBooks() (map[string]booksmodels.Book, error) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	filePath := filepath.Join(pm.dataDir, booksFile)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return make(map[string]booksmodels.Book), nil
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading books file: %v", err)
	}

	var books map[string]booksmodels.Book
	if err := json.Unmarshal(data, &books); err != nil {
		return nil, fmt.Errorf("error unmarshaling books: %v", err)
	}

	log.Printf("Books loaded from %s", filePath)
	return books, nil
}

func (pm *PersistenceManager) SaveOrders(orders map[string]opModels.Order) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	data, err := json.Marshal(orders)
	if err != nil {
		return fmt.Errorf("error marshaling orders: %v", err)
	}

	filePath := filepath.Join(pm.dataDir, ordersFile)
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("error writing orders file: %v", err)
	}

	log.Printf("Orders saved to %s", filePath)
	return nil
}

func (pm *PersistenceManager) LoadOrders() (map[string]opModels.Order, error) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	filePath := filepath.Join(pm.dataDir, ordersFile)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return make(map[string]opModels.Order), nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading orders file: %v", err)
	}

	var orders map[string]opModels.Order
	if err := json.Unmarshal(data, &orders); err != nil {
		return nil, fmt.Errorf("error unmarshaling orders: %v", err)
	}

	log.Printf("Orders loaded from %s", filePath)
	return orders, nil
}
