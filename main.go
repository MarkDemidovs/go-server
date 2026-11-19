package main

import (
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	log.Fatal(srv.ListenAndServe())
}
