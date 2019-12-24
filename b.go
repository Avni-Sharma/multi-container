package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func secHandler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://127.0.0.1:80/c")
	if err != nil {
		log.Fatal("Error from client get", err)
	}
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error from ioutil", err)
	}
	fmt.Fprintf(w, string(resp))
}

func main() {
	log.Print("microservice B started")

	http.HandleFunc("/b", secHandler)

	port := os.Getenv("SVC_PORT")
	if port == "" {
		port = "8222"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
