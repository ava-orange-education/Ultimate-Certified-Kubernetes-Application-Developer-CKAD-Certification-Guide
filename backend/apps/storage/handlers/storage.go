package handlers

import (
	storageRepo "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/repository"
)

type StorageHandler struct {
	br *storageRepo.BooksRepo
	or *storageRepo.OrderRepository
}

func NewStorageHandler(
	br *storageRepo.BooksRepo,
	or *storageRepo.OrderRepository) *StorageHandler {
	return &StorageHandler{br: br, or: or}
}
