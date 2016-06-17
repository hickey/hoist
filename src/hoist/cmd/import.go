package cmd

import "fmt"
import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use:   "import NAME",
	Short: "Read container settings from a file.",
	Run: func(cmd *cobra.Command, args []string) {
		importCommand(args)
	},
}

func importCommand(args []string) {
	fmt.Println("import command: ", args)
}
