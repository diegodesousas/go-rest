package main

import (
	"github.com/diegodesousas/go-rest/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Index).Methods("GET")

	server := &http.Server{
		Handler: r,
		Addr: ":9000",
	}

	log.Fatal(server.ListenAndServe())
}
