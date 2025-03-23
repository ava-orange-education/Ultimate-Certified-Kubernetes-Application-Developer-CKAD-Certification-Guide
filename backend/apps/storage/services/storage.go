package services

import (
	booksmodels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/models"
	opModels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/models"
	storagemodels "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/models"
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/repository"
)

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

func (s *StorageService) GetBookByID(bookID string) (booksmodels.Book, bool) {
	return s.br.GetBookByID(bookID)
}

func (s *StorageService) GetBooks() []booksmodels.Book {
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
