package cmd // Declares the package name as 'cmd'

import (
	"fmt"       // Imports the fmt package for formatted I/O
	"todo/todo" // Imports the todo package from the local project

	"github.com/spf13/cobra" // Imports the cobra package for creating CLI commands
)

// Declares a new cobra.Command struct for the 'list' command
var listCmd = &cobra.Command{
	Use:     "list",              // Command name
	Aliases: []string{"ls", "l"}, // Command aliases
	Short:   "list all tasks",    // Short description of the command
	Long:    `list all tasks`,    // Long description of the command
	Run: func(cmd *cobra.Command, args []string) { // Function to execute when the command is called
		store := &todo.Store{FilePath: "todo.json"} // Creates a new Store instance with the file path 'todo.json'
		tasks, err := store.LoadTasks()             // Loads tasks from the file
		if err != nil {                             // Checks if there was an error loading tasks
			fmt.Println("Error loading tasks: ", err) // Prints the error message
			return                                    // Exits the function if there was an error
		}

		if len(tasks) == 0 { // Checks if there are no tasks
			fmt.Println("No tasks to list") // Prints a message indicating no tasks
			return                          // Exits the function if there are no tasks
		}

		// Create a map to group tasks by their group
		groupTasks := make(map[string][]todo.Task) // Creates a map to store tasks by group
		for _, task := range tasks {
			// fmt.Println(task)                                             // Iterates over the tasks
			groupTasks[task.Group] = append(groupTasks[task.Group], task) // Appends the task to the group
		}

		fmt.Println("List All Tasks:") // Prints a header for the task list
		// for i, task := range tasks { // Iterates over the tasks
		// 	fmt.Printf("%d. %s\n", i+1, task.Description) // Prints each task with its index
		// }

		// Print tasks by group
		for group, tasks := range groupTasks {
			fmt.Printf(">>> Group : %s\n", group)
			for i, task := range tasks {
				status := "Pending"
				if task.Status == todo.Completed {
					status = "Completed"
				}
				fmt.Printf("%d. %s\t status: %s\n", i+1, task.Description, status)
			}
		}
	},
}

func init() { // init function is called when the package is initialized
	rootCmd.AddCommand(listCmd) // Adds the 'list' command to the root command
}
