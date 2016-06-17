package cmd

import "fmt"
import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the current container startup status.",
	Run: func(cmd *cobra.Command, args []string) {
		statusCommand()
	},
}

func statusCommand() {
	fmt.Println("status command")
}
