package cmd

import (
	"fmt"
	"strings"
	"todo/todo"

	"github.com/spf13/cobra"
)

var deleteDescription string

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "d"},
	Short:   "Delete a task from the list",
	Long:    `Delete a task from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		// 删除任务 必须输入 -d 参数，匹配任务的描述
		if deleteDescription == "" {
			fmt.Println("Please provide a task description to delete")
			return
		}

		store := &todo.Store{FilePath: "todo.json"}
		tasks, err := store.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks: ", err)
			return
		}

		var updatedTasks []todo.Task
		taskFound := false

		for _, task := range tasks {
			if strings.TrimSpace(task.Description) != strings.TrimSpace(deleteDescription) {
				updatedTasks = append(updatedTasks, task)
			} else {
				taskFound = true
			}
		}

		if !taskFound {
			fmt.Printf("Task with description '%s' not found\n", deleteDescription)
			return
		}

		err = store.SaveTasks(updatedTasks)

		if err != nil {
			fmt.Println("Error saving tasks: ", err)
			return
		}

		fmt.Printf("Deleted task: %s\n", deleteDescription)

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&deleteDescription, "description", "d", "", "Task description")
	deleteCmd.MarkFlagRequired("description")
}
