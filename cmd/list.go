package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "list all tasks",
	Long:    `list all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No tasks to list")
			return
		}

		fmt.Println("List Tasks")
		for i, task := range args {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
