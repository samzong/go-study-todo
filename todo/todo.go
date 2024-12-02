// Package todo defines a simple task management system
package todo

// Task represents a single task with a description
type Task struct {
	// Description is a brief summary of the task
	// The `json:"description"` tag specifies that this field should be serialized to JSON with the key "description"
	Description string `json:"description"`
}
