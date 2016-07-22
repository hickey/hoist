package cmd

import (
	"app"
	"fmt"
	"github.com/spf13/cobra"
)

var label string

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().StringVarP(&label, "label", "l", "LABEL", "Add a label")
}

var addCmd = &cobra.Command{
	Use:   "add NAME",
	Short: "Add a Docker container to be managed.",
	Run: func(cmd *cobra.Command, args []string) {
		addCommand(args)
	},
}

func addCommand(args []string) {
	fmt.Println("add command: ", args)
	fmt.Println("labels: ", label)
	fmt.Println("verbose: ", app.GetSettingBool("verbose"))
}
