package cmd

import (
	"fmt"
	"todo/todo"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "list all tasks",
	Long:    `list all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		// 从文件中读取任务列表，并打印
		store := &todo.Store{FilePath: "todo.json"}
		tasks, err := store.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks: ", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks to list")
			return
		}

		fmt.Println("List Tasks")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
