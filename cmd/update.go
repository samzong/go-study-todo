package cmd // Declares the package name

import (
	"fmt"       // Imports the fmt package for formatted I/O
	"strings"   // Imports the strings package for string manipulation
	"todo/todo" // Imports the todo package from the local project

	"github.com/spf13/cobra" // Imports the cobra package for creating CLI applications
)

var (
	message    string // Variable to store the task description
	status     string // Variable to store the task status
	group      string // Variable to store the task group
	newMessage string // Variable to store the new task description
)

// updateCmd represents the 'update' command in the CLI application
var updateCmd = &cobra.Command{
	Use:     "update",               // Command name
	Aliases: []string{"u"},          // Alternative command names
	Short:   "Update a task config", // Short description of the command
	Long:    `Update a task config`, // Long description shown in help
	Run: func(cmd *cobra.Command, args []string) { // Function to execute when the command is called
		// Get the value of the 'description' flag from the command
		message, _ = cmd.Flags().GetString("description")
		// Get the value of the 'status' flag from the command
		status, _ = cmd.Flags().GetString("status")
		// Get the value of the 'group' flag from the command
		group, _ = cmd.Flags().GetString("group")
		// Get the value of the 'new-description' flag from the command
		newMessage, _ = cmd.Flags().GetString("new-description")

		// Create a new Store instance with the file path 'todo.json'
		store := &todo.Store{FilePath: "todo.json"}
		// Load tasks from the store
		tasks, err := store.LoadTasks()
		if err != nil { // Check for errors
			fmt.Println("Error loading tasks: ", err) // Print error message
			return                                    // Exit the function
		}

		if len(tasks) == 0 { // Check if there are no tasks
			fmt.Println("No tasks to update") // Print message
		}

		// Iterate over the tasks
		for i, task := range tasks {
			// Check if the task description matches the input description
			if strings.TrimSpace(task.Description) == strings.TrimSpace(message) {
				if newMessage != "" { // Check if a new description is provided
					tasks[i].Description = newMessage // Update the task description
				}
				if status != "" { // Check if a status is provided
					switch strings.ToLower(status) { // Convert status to lowercase and switch
					case "completed":
						tasks[i].Status = todo.Completed // Set status to completed
					case "pending":
						tasks[i].Status = todo.Pending // Set status to pending
					default:
						fmt.Println("Invalid status") // Print invalid status message
						return                        // Exit the function
					}
				}
				if group != "" { // Check if a group is provided
					tasks[i].Group = group // Update the task group
				}

				// Save the updated tasks to the store
				err := store.SaveTasks(tasks)
				if err != nil { // Check for errors
					fmt.Println("Error saving tasks: ", err) // Print error message
					return                                   // Exit the function
				}
				fmt.Println("Task updated successfully") // Print success message
			}
		}
	},
}

func init() {
	// Register the 'update' command with the root command
	rootCmd.AddCommand(updateCmd)
	// Register the 'description' flag with the 'update' command
	updateCmd.Flags().StringVarP(&message, "description", "d", "", "Task description <required>")
	updateCmd.MarkFlagRequired("description") // Mark the 'description' flag as required
	// Register the 'new-description' flag with the 'update' command
	updateCmd.Flags().StringVarP(&newMessage, "new-description", "n", "", "New Task Description")
	// Register the 'group' flag with the 'update' command
	updateCmd.Flags().StringP("group", "g", "default", "Task group")
	// Register the 'status' flag with the 'update' command
	updateCmd.Flags().StringP("status", "s", "pending", "Task status")
}
