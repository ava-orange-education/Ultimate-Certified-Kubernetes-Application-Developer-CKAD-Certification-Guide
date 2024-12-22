package handlers

import (
	storageRepo "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/storage/repository"
)

type StorageHandler struct {
	br *storageRepo.BooksRepo
}

func NewStorageHandler(br *storageRepo.BooksRepo) *StorageHandler {
	return &StorageHandler{br: br}
}
