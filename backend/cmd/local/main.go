package main

import (
	"io"
	"log"
	"net/http"

	"github.com/m4tthewde/happening/backend/internal"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == "OPTIONS" {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	response, status := internal.HandleNewSubscription(string(body))
	w.WriteHeader(status)
	w.Write([]byte(response))

}

func main() {
	http.HandleFunc("/api", HandleRequest)

	log.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}

type NewSubscriptionBody struct {
	TargetUserID string `json:"target_id"`
	SubType      string `json:"subscription_type"`
}
