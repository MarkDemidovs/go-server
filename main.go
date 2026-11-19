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

	router.Handle("/", http.FileServer(http.Dir(".")))

	log.Fatal(srv.ListenAndServe())
}
