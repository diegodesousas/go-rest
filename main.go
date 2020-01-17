package main

import (
	"fmt"
	"github.com/diegodesousas/go-rest/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Index).Methods("GET")
	r.HandleFunc("/hello-world", handler.HelloWorld).Methods("POST")

	server := &http.Server{
		Handler: r,
		Addr: ":9000",
	}

	fmt.Println("Running on port 9000")
	log.Fatal(server.ListenAndServe())
}
