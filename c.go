package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func thirdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "response from C microservice")
}

func main() {
	log.Print("microservice C started")

	http.HandleFunc("/c", thirdHandler)

	port := os.Getenv("SVC_PORT")
	if port == "" {
		port = "8333"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
