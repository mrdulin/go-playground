package main

import (
	"fmt"
	cloudrun "go-clean-arch/infrastructure/cloudrun/reports"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/google/reports/ad", cloudrun.GetAdPerformanceReport)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
