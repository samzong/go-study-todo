package api

import (
	"log"
	"net/http"
)

func todosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	log.Println(r.Method, "/api/v1/todos")

	switch r.Method {
	case http.MethodGet:
		handleGetTodos(w, r)
	case http.MethodPost:
		handlePostTodos(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
