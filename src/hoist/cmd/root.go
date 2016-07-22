package cmd

import (
	//"app"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Turn on verbose logging")
	RootCmd.PersistentFlags().BoolP("debug", "d", false, "Print debugging logs")
	viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))
}

func processSwitches() {
	fmt.Println("debug = ", viper.Get("debug"))
}
