package todo

// Importing necessary packages.
// "encoding/json" is used for encoding and decoding JSON data.
// "os" provides a platform-independent interface to operating system functionality, such as file handling.
import (
	"encoding/json"
	"os"
)

// Store struct represents a storage for tasks.
// It contains a single field FilePath which is a string indicating the file location.
type Store struct {
	FilePath string // FilePath is the path to the file where tasks are stored.
}

// LoadTasks method loads tasks from the file specified in Store.FilePath.
// It returns a slice of Task structs and an error if any occurs during file operations.
func (s *Store) LoadTasks() ([]Task, error) {
	// Open the file located at FilePath for reading.
	file, err := os.Open(s.FilePath)
	// Check if there was an error opening the file.
	if err != nil {
		// If the error is that the file does not exist, return an empty slice of tasks and no error.
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		// Return nil and the error if any other error occurred.
		return nil, err
	}
	// Ensure the file is closed after the function completes.
	defer file.Close()

	// Declare a variable to hold the tasks loaded from the file.
	var Tasks []Task
	// Decode the JSON data from the file into the Tasks slice.
	err = json.NewDecoder(file).Decode(&Tasks)
	// If decoding fails, return nil and the error.
	if err != nil {
		return nil, err
	}

	// Return the loaded tasks and no error.
	return Tasks, nil
}

// SaveTasks method saves the provided tasks to the file specified in Store.FilePath.
// It takes a slice of Task structs and returns an error if any occurs during file operations.
func (s *Store) SaveTasks(tasks []Task) error {
	// Marshal the tasks slice into JSON format.
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return os.WriteFile(s.FilePath, data, 0644)
}
