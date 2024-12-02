package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo/todo"

	"github.com/spf13/cobra"
)

var message string

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create", "new", "n", "a"},
	Short:   "Add a new task to the list",
	Long:    `Add a new task to the list`,
	Run: func(cmd *cobra.Command, args []string) {
		var todoDescription string
		// 从键盘输入任务，并保存到文件

		if message != "" {
			todoDescription = message
		} else {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter the task to add: ")
			todoDescription, _ = reader.ReadString('\n')
			todoDescription = strings.TrimSpace(todoDescription)
		}

		if todoDescription == "" {
			fmt.Println("Please provide a task to add")
			return
		}
		store := &todo.Store{FilePath: "todo.json"}
		tasks, err := store.LoadTasks()

		if err != nil {
			fmt.Println("Error loading tasks: ", err)
			return
		}

		tasks = append(tasks, todo.Task{Description: todoDescription})
		err = store.SaveTasks(tasks)
		if err != nil {
			fmt.Println("Error saving task:", err)
			return
		}
		fmt.Printf("Added task: %s\n", todoDescription)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&message, "description", "d", "", "Task description")
}
