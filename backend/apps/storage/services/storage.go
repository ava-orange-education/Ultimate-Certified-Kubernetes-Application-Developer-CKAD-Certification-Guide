package services

import (
	"os"
	"sync"

	booksmodels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/models"
	opModels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/models"
	storagemodels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/models"
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/repository"
)

// FeatureFlags represents feature toggles for the service
type FeatureFlags struct {
	EnableNewCaching      bool
	EnableBatchProcessing bool
	UseNewOrderFormat     bool
}

// GetFeatureFlags returns the current feature flags
func GetFeatureFlags() FeatureFlags {
	// In a real system, this might come from a config service or environment
	version := os.Getenv("SERVICE_VERSION")

	// Enable new features only in canary version
	if version == "canary" {
		return FeatureFlags{
			EnableNewCaching:      true,
			EnableBatchProcessing: true,
			UseNewOrderFormat:     true,
		}
	}

	// Default stable version features
	return FeatureFlags{
		EnableNewCaching:      false,
		EnableBatchProcessing: false,
		UseNewOrderFormat:     false,
	}
}

type StorageService struct {
	br *repository.BooksRepo
	or *repository.OrderRepository
}

func NewStorageService(
	br *repository.BooksRepo,
	or *repository.OrderRepository) *StorageService {
	return &StorageService{
		br: br,
		or: or,
	}
}

// Book operations
func (s *StorageService) AddBook(book booksmodels.Book) {
	s.br.AddBook(book)
}

// Cache for books to demonstrate the new caching feature
var (
	bookCache     = make(map[string]booksmodels.Book)
	bookCacheLock = &sync.RWMutex{}
)

// Helper functions for the new caching feature
func getBookFromCache(id string) (booksmodels.Book, bool) {
	bookCacheLock.RLock()
	defer bookCacheLock.RUnlock()
	book, exists := bookCache[id]
	return book, exists
}

func addBookToCache(id string, book booksmodels.Book) {
	bookCacheLock.Lock()
	defer bookCacheLock.Unlock()
	bookCache[id] = book
}

func (s *StorageService) GetBookByID(bookID string) (booksmodels.Book, bool) {
	flags := GetFeatureFlags()

	// Use different caching strategy based on feature flag
	if flags.EnableNewCaching {
		// New caching strategy implementation
		// In a real system, this might use a more efficient caching mechanism
		// For demonstration, we're implementing a simple in-memory cache check

		// First check if the book is in the cache
		if book, found := getBookFromCache(bookID); found {
			// Book found in cache, return it
			return book, true
		}

		// If not in cache, get from repository and add to cache
		book, exists := s.br.GetBookByID(bookID)
		if exists {
			addBookToCache(bookID, book)
		}
		return book, exists
	}

	// Use existing implementation for stable version (no caching)
	return s.br.GetBookByID(bookID)
}

func (s *StorageService) GetBooks() []booksmodels.Book {
	flags := GetFeatureFlags()

	// Use different retrieval strategy based on feature flag
	if flags.EnableBatchProcessing {
		// New batch processing implementation
		// In a real system, this might use pagination or chunking for better performance
		// For demonstration, we're implementing a simple batched retrieval

		// Get all books
		allBooks := s.br.GetBooks()

		// Process books in batches (simulated)
		// This could represent processing optimization for large datasets
		processedBooks := make([]booksmodels.Book, 0, len(allBooks))
		batchSize := 10

		for i := 0; i < len(allBooks); i += batchSize {
			end := i + batchSize
			if end > len(allBooks) {
				end = len(allBooks)
			}

			// Process this batch (in a real system, this would do something meaningful)
			batch := allBooks[i:end]

			// Add processed batch to result
			processedBooks = append(processedBooks, batch...)
		}

		return processedBooks
	}

	// Use existing implementation for stable version (no batching)
	return s.br.GetBooks()
}

func (s *StorageService) UpdateBook(book booksmodels.Book) error {
	return s.br.UpdateBook(book)
}

func (s *StorageService) CheckQuantity(bookID string) (int, bool) {
	book, exists := s.br.GetBookByID(bookID)
	if !exists {
		return 0, false
	}
	return book.Quantity, true
}

func (s *StorageService) UpdateBookQuantity(updateReq storagemodels.UpdateBookQuantityRequest) (booksmodels.Book, error) {
	return s.br.UpdateBookQuantity(updateReq)
}

// Order operations
func (s *StorageService) ListOrders() []opModels.Order {
	flags := GetFeatureFlags()

	// Use different order format based on feature flag
	if flags.UseNewOrderFormat {
		// New order format implementation
		// In a real system, this might use a different data structure or include additional fields
		// For demonstration, we're adding a version field to each order

		// Get orders using existing implementation
		orders := s.or.ListOrders()

		// Enhance orders with additional information (simulating new format)
		for i := range orders {
			// Add a version indicator to the order ID to simulate format change
			// In a real system, this might involve more substantial changes to the data structure
			if orders[i].ID != "" {
				orders[i].ID = orders[i].ID + "-v2"
			}

			// Could also add new fields or transform existing ones in a real implementation
		}

		return orders
	}

	// Use existing implementation for stable version (original format)
	return s.or.ListOrders()
}

func (s *StorageService) AddOrder(order opModels.Order) {
	s.or.AddOrder(order)
}

func (s *StorageService) GetOrderByID(orderID string) (opModels.Order, bool) {
	return s.or.GetOrderByID(orderID)
}

func (s *StorageService) UpdateOrderStatus(updateReq opModels.UpdateOrderStatusRequest) (opModels.Order, error) {
	return s.or.UpdateOrderStatus(updateReq)
}
