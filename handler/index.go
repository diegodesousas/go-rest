package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request)  {
	var h HelloRequest
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &h)
	if err != nil {
		log.Fatal(err)
	}

	resp := HelloResponse{
		Message: fmt.Sprintf("Hello, %s", h.Name),
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Fatal(err)
	}
}