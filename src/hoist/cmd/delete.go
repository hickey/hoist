package cmd

import "fmt"
import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete NAME",
	Short: "Remove a Docker container being managed.",
	Run: func(cmd *cobra.Command, args []string) {
		deleteCommand(args)
	},
}

func deleteCommand(args []string) {
	fmt.Println("delete command: ", args)
}
