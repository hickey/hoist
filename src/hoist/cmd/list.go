package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"hoist/store"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Docker containers being managed",
	Run: func(cmd *cobra.Command, args []string) {
		listCommand()
	},
}

func listCommand() {
	fmt.Println("Current list of managed containers:")
	for _, container := range store.StartupList() {
		fmt.Printf("\t%s\n", container)
	}
}
