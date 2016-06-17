package cmd

import "fmt"
import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(exportCmd)
}

var exportCmd = &cobra.Command{
	Use:   "export NAME",
	Short: "Export a container settings to a file.",
	Run: func(cmd *cobra.Command, args []string) {
		exportCommand(args)
	},
}

func exportCommand(args []string) {
	fmt.Println("export command: ", args)
}
