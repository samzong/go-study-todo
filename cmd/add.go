package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create", "new", "n", "a"},
	Short:   "Add a new task to the list",
	Long:    `Add a new task to the list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add Item: ", args)
		if len(args) < 1 {
			fmt.Println("Please specify the task to add")
			return
		}
		todo := args[0]
		fmt.Printf("Added task: %s\n", todo)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
