package cmd

import (
	"fmt"       // Imports the fmt package for formatted I/O
	"todo/todo" // Imports the todo package from the local project

	"github.com/spf13/cobra" // Imports the cobra package for creating CLI commands
)

// Declares a new cobra.Command struct for the 'list' command
var doneCmd = &cobra.Command{
	Use: "done", // Command name
	// Aliases: []string{"ls", "l"}, // Command aliases
	Short: "list all done tasks", // Short description of the command
	Long:  `list all done tasks`, // Long description of the command
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PreRun done")
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentPreRun done")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PostRun done")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentPostRun done")
	},
	Run: func(cmd *cobra.Command, args []string) {
		store := &todo.Store{FilePath: "todo.json"} // Creates a new Store instance with the file path 'todo.json'
		tasks, err := store.LoadTasks()             // Loads tasks from the file

		if err != nil {
			fmt.Println("Error loading tasks: ", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks to list")
			return
		}

		groupedTasks := make(map[string][]todo.Task) // Create a map to group tasks by their group
		for _, task := range tasks {
			if task.Status == todo.Completed {
				// fmt.Printf("Task: %s\t status: %s\n", task.Description, task.Status)
				groupedTasks[task.Group] = append(groupedTasks[task.Group], task)
			}
		}

		// Print completed tasks by group
		fmt.Println("List Completed Tasks:")
		for group, tasks := range groupedTasks {
			fmt.Printf(">>> Group: %s\n", group)
			for _, task := range tasks {
				fmt.Printf("Task: %s\tStatus: %s\n", task.Description, task.Status)
			}
		}
	},
}

func init() { // init function is called when the package is initialized
	listCmd.AddCommand(doneCmd) // Adds the 'list' command to the root command
}
