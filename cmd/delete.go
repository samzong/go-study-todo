package cmd // Define the package name as 'cmd'

import (
	"fmt"       // Import the fmt package for formatted I/O
	"strings"   // Import the strings package for string manipulation
	"todo/todo" // Import the todo package from the local project

	"github.com/spf13/cobra" // Import the cobra package for creating CLI commands
)

var deleteDescription string // Declare a variable to hold the task description to delete

// Define the delete command using cobra.Command struct
var deleteCmd = &cobra.Command{
	Use:     "delete",                      // Command name
	Aliases: []string{"del", "d"},          // Command aliases
	Short:   "Delete a task from the list", // Short description of the command
	Long:    `Delete a task from the list`, // Long description of the command
	Run: func(cmd *cobra.Command, args []string) { // Function to execute when the command is called
		// Check if the deleteDescription is provided
		if deleteDescription == "" {
			fmt.Println("Please provide a task description to delete") // Print error message if not provided
			return                                                     // Exit the function
		}

		// Create a new Store instance with the file path 'todo.json'
		store := &todo.Store{FilePath: "todo.json"}
		// Load tasks from the store
		tasks, err := store.LoadTasks()
		if err != nil { // Check for errors
			fmt.Println("Error loading tasks: ", err) // Print error message
			return                                    // Exit the function
		}

		var updatedTasks []todo.Task // Declare a slice to hold the updated tasks
		taskFound := false           // Flag to check if the task is found

		// Iterate over the tasks
		for _, task := range tasks {
			// Compare task descriptions, ignoring leading/trailing spaces
			if strings.TrimSpace(task.Description) != strings.TrimSpace(deleteDescription) {
				updatedTasks = append(updatedTasks, task) // Add task to updatedTasks if it doesn't match
			} else {
				taskFound = true // Set flag to true if task is found
			}
		}

		if !taskFound { // Check if the task was not found
			fmt.Printf("Task with description '%s' not found\n", deleteDescription) // Print error message
			return                                                                  // Exit the function
		}

		// Save the updated tasks back to the store
		err = store.SaveTasks(updatedTasks)
		if err != nil { // Check for errors
			fmt.Println("Error saving tasks: ", err) // Print error message
			return                                   // Exit the function
		}

		// Print success message
		fmt.Printf("Deleted task: %s\n", deleteDescription)
	},
}

func init() { // init function to initialize the command
	rootCmd.AddCommand(deleteCmd) // Add deleteCmd to the root command
	// Define a string flag for the delete command
	deleteCmd.Flags().StringVarP(&deleteDescription, "description", "d", "", "Task description") // Define the 'description' flag with shorthand 'd'
	deleteCmd.MarkFlagRequired("description")                                                    // Mark the description flag as required
}
