package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"todo/todo"
)

func handlePostTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var task todo.Task
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	log.Printf("Received todo: %+v", task)

	store := &todo.Store{FilePath: "todo.json"}
	// Load existing tasks from the file
	tasks, err := store.LoadTasks()
	// Handle any errors that occur during task loading
	if err != nil {
		log.Println("Error loading tasks: ", err) // Output the error message
		return                                    // Exit the function
	}

	// Check the task exists
	if len(tasks) != 0 {
		for _, t := range tasks {
			if strings.TrimSpace(t.Description) == strings.TrimSpace(task.Description) {
				log.Printf("Task already exists")
				http.Error(w, "Task already exists", http.StatusBadRequest)
				return
			}
		}
	}

	// Append the new task to the existing list of tasks
	tasks = append(tasks, todo.Task{Description: task.Description, Group: task.Group, Status: todo.Pending})

	// Save the updated list of tasks back to the file
	err = store.SaveTasks(tasks)
	// Handle any errors that occur during task saving
	if err != nil {
		log.Println("Error saving task:", err) // Output the error message
		http.Error(w, "Error saving task", http.StatusBadRequest)
		return // Exit the function
	}

	// Confirm to the user that the task has been added successfully
	log.Printf("Added task: %s, Group is: %s", task.Description, task.Group)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	// repsone task add successfully
	response := map[string]string{"message": "Task added successfully"}
	json.NewEncoder(w).Encode(response)
}
