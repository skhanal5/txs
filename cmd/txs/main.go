package main

import (
	"log"
	"net/http"

	"github.com/skhanal5/txs/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.GetHealth)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}