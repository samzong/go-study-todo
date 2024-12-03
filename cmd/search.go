/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"todo/todo"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		store := &todo.Store{FilePath: "todo.json"} // Creates a new Store instance with the file path 'todo.json'
		tasks, err := store.LoadTasks()             // Loads tasks from the file

		if err != nil {
			panic(err)
			fmt.Println("Error loading tasks: ", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks to list")
			return
		}

		groupedTasks := make(map[string][]todo.Task) // Create a map to group tasks by their group

		group := strings.TrimSpace(cmd.Flag("group").Value.String())
		keywords := strings.TrimSpace(cmd.Flag("keywords").Value.String())

		if keywords == "" {
			fmt.Println("Please provide a keywords")
			return
		}

		if group != "" {
			for _, task := range tasks {
				if task.Group == group {
					if strings.Contains(task.Description, keywords) {
						groupedTasks[task.Group] = append(groupedTasks[task.Group], task)
					}
				}
			}
		} else {
			for _, task := range tasks {
				if strings.Contains(task.Description, keywords) {
					groupedTasks[task.Group] = append(groupedTasks[task.Group], task)
				}
			}
		}

		if len(groupedTasks) == 0 {
			fmt.Println("No tasks found")
			return
		}

		// Print completed tasks by group
		for group, tasks := range groupedTasks {
			fmt.Printf(">>> Group: %s\n", group)
			for _, task := range tasks {
				fmt.Printf("Task: %s\tStatus: %s\n", task.Description, task.Status)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("keywords", "k", "", "Search tasks by keywords <required>")
	searchCmd.MarkFlagRequired("keywords")
	searchCmd.Flags().StringP("group", "g", "", "Group tasks by group")
}
