package main

import (
	"net/http"

	"github.com/ras0q/go-wstest/ws"
)

func main() {
	http.HandleFunc("/ws", ws.Serve)
	http.ListenAndServe(":8080", nil)
}
