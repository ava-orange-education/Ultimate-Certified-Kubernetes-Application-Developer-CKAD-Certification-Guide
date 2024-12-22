package main

import (
	"log"
	"net/http"

	opRepo "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/repository"
	opSvc "github.com/ava-orange-education/Ultimate-Certified-Kubernetes-Application-Developer-CKAD-Certification-Guide/backend/apps/order-processor/services"
)

const (
	StorageServiceURL          = "http://storage-service:8083"
	OrderProcessingServicePort = "8082"
)

func main() {
	opr := opRepo.NewOrderRepository()
	ops := opSvc.NewOrderProcessingService(opr)

	router := ops.AddRoutes()

	log.Printf("Order Processor Service running on :%s", OrderProcessingServicePort)
	if err := http.ListenAndServe(":"+OrderProcessingServicePort, router); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
