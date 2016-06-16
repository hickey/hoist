package cmd

import "fmt"
import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Docker containers being managed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list command")
	},
}
