package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexBody struct {
	Project string `json:"project"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	body := IndexBody{
		Project: "Challenge Bravo Golang",
	}

	response, err := json.Marshal(body)
	if err != nil {
		log.Fatal("error when marshall json")
	}

	_, err = w.Write(response)
	if err != nil {
		log.Fatal("error when write response body")
	}
}
