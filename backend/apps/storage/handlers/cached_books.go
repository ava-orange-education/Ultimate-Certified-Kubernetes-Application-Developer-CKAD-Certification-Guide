package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/repository"
	"github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/services"
	httpPkg "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/pkg/http"
	"github.com/go-chi/chi/v5"
)

type CachedBooksHandler struct {
	service *services.StorageService
	cache   *repository.CacheManager
}

func NewCachedBooksHandler(service *services.StorageService) *CachedBooksHandler {
	// Create cache with 5 minute TTL
	cache := repository.NewCacheManager(5 * time.Minute)

	return &CachedBooksHandler{
		service: service,
		cache:   cache,
	}
}

func (cbh *CachedBooksHandler) GetBookDetails(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "id")
	if bookID == "" {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var bookResponse map[string]any
	found, err := cbh.cache.Get("book_"+bookID, &bookResponse)
	if err != nil {
		log.Printf("Cache error: %v", err)
	}

	if found {
		log.Printf("Cache hit for book %s", bookID)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Cache", "HIT")
		json.NewEncoder(w).Encode(bookResponse)
		return
	}

	book, exists := cbh.service.GetBookByID(bookID)
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	if err := cbh.cache.Set("book_"+bookID, book); err != nil {
		log.Printf("Failed to cache book(%s): %v", bookID, err)
	}

	w.Header().Set("X-Cache", "MISS")
	httpPkg.JSON(r.Context(), w, http.StatusOK, book)
}

func (cbh *CachedBooksHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var booksResponse map[string]any

	found, err := cbh.cache.Get("all_books", &booksResponse)
	if err != nil {
		log.Printf("Cache error: %v", err)
	}

	if found {
		log.Printf("Cache hit for all books")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Cache", "HIT")
		json.NewEncoder(w).Encode(booksResponse)
		return
	}

	books := cbh.service.GetBooks()

	if err := cbh.cache.Set("all_books", books); err != nil {
		log.Printf("Failed to cache books: %v", err)
	}

	w.Header().Set("X-Cache", "MISS")
	httpPkg.JSON(r.Context(), w, http.StatusOK, books)
}

func (cbh *CachedBooksHandler) AddCachedRoutes(router *chi.Mux) {
	router.Get("/cached-books/{id}", cbh.GetBookDetails)
	router.Get("/cached-books", cbh.GetAllBooks)
}
