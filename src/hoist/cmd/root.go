package cmd

import (
	//"app"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "hoist",
	Short: "hoist is a tool to automate the starting of Docker containers",
	Long:  `foo`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		processSwitches()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var verbose bool
var debug bool

func init() {
	cobra.OnInitialize()
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Turn on verbose logging")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Print debugging logs")
}

func processSwitches() {

}
