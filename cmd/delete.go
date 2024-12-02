package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "d"},
	Short:   "Delete a task from the list",
	Long:    `Delete a task from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete Item: ", args)
		if len(args) < 1 {
			fmt.Println("Please specify the task to delete")
			return
		}
		todo := args[0]
		fmt.Printf("Deleted task: %s\n", todo)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
