package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(GreetingServer)

	err := http.ListenAndServe(":8000", handler)

	if err != nil {
		log.Fatalf("erro ao abrir a porta 8000 %v", err)
	}
}