package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "hoist",
	Short: "hoist is a tool to automate the starting of Docker containers",
	Long:  `foo`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	cobra.OnInitialize()
}
