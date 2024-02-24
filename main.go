package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Greeting struct {
	Ok       bool   `json:"ok"`
	Greeting string `json:"greeting"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	log.Printf("Received name=%s", name)
	greeting := Greeting{
		Ok:       true,
		Greeting: fmt.Sprintf("Hello %s", name),
	}
	body, err := json.Marshal(&greeting)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to marshal response: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Failed to write response %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /echo/{name}", echoHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
