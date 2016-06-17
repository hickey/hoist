package cmd

import "fmt"
import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(modifyCmd)
}

var modifyCmd = &cobra.Command{
	Use:   "modify NAME",
	Short: "Change a container's settings.",
	Run: func(cmd *cobra.Command, args []string) {
		modifyCommand(args)
	},
}

func modifyCommand(args []string) {
	fmt.Println("modify command: ", args)
}
