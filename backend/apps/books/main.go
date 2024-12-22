package main

import (
	"log"
	"net/http"

	svcs "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/books/services"
)

const BooksServicePort = "8081"

func main() {
	service := svcs.NewBooksService()

	router := service.AddRoutes()

	log.Printf("Books Service running on :%s", BooksServicePort)
	if err := http.ListenAndServe(":"+BooksServicePort, router); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
