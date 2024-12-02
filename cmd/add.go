// Package cmd defines the commands for the CLI application
package cmd

import (
	// bufio provides buffered I/O operations
	"bufio"
	// fmt implements formatted I/O functions
	"fmt"
	// os provides a platform-independent interface to operating system functionality
	"os"
	// strings contains functions to manipulate UTF-8 encoded strings
	"strings"
	// todo is a custom package for task management
	"todo/todo"

	// cobra is a library for creating powerful modern CLI applications
	"github.com/spf13/cobra"
)

// message holds the task description provided via command-line flag
var message string

// addCmd represents the 'add' command in the CLI application
var addCmd = &cobra.Command{
	Use:     "add",                               // Command name
	Aliases: []string{"create", "new", "n", "a"}, // Alternative command names
	Short:   "Add a new task to the list",        // Short description of the command
	Long:    `Add a new task to the list`,        // Long description shown in help
	Run: func(cmd *cobra.Command, args []string) { // Function to execute when the command is called
		var todoDescription string // Variable to store the task description

		// Check if the 'message' flag is provided
		if message != "" {
			todoDescription = message // Use the provided message as the task description
		} else {
			// If no message is provided, prompt the user for input
			reader := bufio.NewReader(os.Stdin)                  // Create a new buffered reader from standard input
			fmt.Print("Enter the task to add: ")                 // Prompt message
			todoDescription, _ = reader.ReadString('\n')         // Read input until a newline character
			todoDescription = strings.TrimSpace(todoDescription) // Remove leading and trailing whitespace
		}

		// Validate that a task description is provided
		if todoDescription == "" {
			fmt.Println("Please provide a task to add") // Inform the user that input is required
			return                                      // Exit the function
		}

		// Create a new store instance with the specified file path
		store := &todo.Store{FilePath: "todo.json"}
		// Load existing tasks from the file
		tasks, err := store.LoadTasks()
		// Handle any errors that occur during task loading
		if err != nil {
			fmt.Println("Error loading tasks: ", err) // Output the error message
			return                                    // Exit the function
		}

		// Append the new task to the existing list of tasks
		tasks = append(tasks, todo.Task{Description: todoDescription})
		// Save the updated list of tasks back to the file
		err = store.SaveTasks(tasks)
		// Handle any errors that occur during task saving
		if err != nil {
			fmt.Println("Error saving task:", err) // Output the error message
			return                                 // Exit the function
		}

		// Confirm to the user that the task has been added successfully
		fmt.Printf("Added task: %s\n", todoDescription)
	},
}

// init initializes the command and adds it to the root command
func init() {
	rootCmd.AddCommand(addCmd)                                                      // Register the addCmd with the root command
	addCmd.Flags().StringVarP(&message, "description", "d", "", "Task description") // Define the 'description' flag with shorthand 'd'
}
