package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/didip/tollbooth/v7"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	message := Message{
		Status: "successful",
		Body:   "Hi! You are reached the API, How mat i help you?",
	}

	err := json.NewEncoder(w).Encode(&message)
	if err != nil {
		return
	}
}

func main() {
	message := Message{
		Status: "Request Failed.",
		Body:   "The API is at capacity, try again later",
	}

	jsonMessag, _ := json.Marshal(message)
	tlbthLimiter := tollbooth.NewLimiter(1, nil)
	tlbthLimiter.SetMessageContentType("application/json")
	tlbthLimiter.SetMessage(string(jsonMessag))

	http.Handle("/ping", tollbooth.LimitFuncHandler(tlbthLimiter, endpointHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("There was an error listening on port :8080", err)
	}
}
