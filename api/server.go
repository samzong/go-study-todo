package api

import (
	"log"
	"net/http"
)

func StartServer(address string, port string) {
	log.Println("Starting server on :8080")
	addr := address + ":" + port
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

func init() {
	http.HandleFunc("/", handler)
}
