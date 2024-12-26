package main

import (
	"log"
	"net/http"

	opSvc "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/services"
)

const (
	OrderProcessingServicePort = "8082"
)

func main() {
	ops := opSvc.NewOrderProcessingService()

	router := ops.AddRoutes()

	log.Printf("Order Processor Service running on :%s", OrderProcessingServicePort)
	if err := http.ListenAndServe(":"+OrderProcessingServicePort, router); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
