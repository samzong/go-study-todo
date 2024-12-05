package api

import (
	"encoding/json"
	"net/http"
	"todo/todo"
)

func handleGetTodos(w http.ResponseWriter, r *http.Request) {
	store := &todo.Store{FilePath: "todo.json"}
	tasks, err := store.LoadTasks()
	if err != nil {
		http.Error(w, "Error loading tasks", http.StatusInternalServerError)
		return
	}
	if len(tasks) == 0 {
		http.Error(w, "No tasks to list", http.StatusNotFound)
		return
	}

	groupTasks := make(map[string][]todo.Task)
	for _, task := range tasks {
		groupTasks[task.Group] = append(groupTasks[task.Group], task)
	}
	json.NewEncoder(w).Encode(groupTasks)
}
