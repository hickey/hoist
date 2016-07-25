package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*

  hoist add --env VAR --label LAB=VAL --container FOO --volume DIR=DIR --port PORT=PORT --link MACH=MACH BAR

*/

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().BoolP("label", "l", false, "Add a label")
	addCmd.PersistentFlags().BoolP("env", "e", false, "Add environmental variable")
	addCmd.PersistentFlags().BoolP("volume", "V", false, "Add volume mapping")
	addCmd.PersistentFlags().BoolP("port", "p", false, "Docker port mapping")
	addCmd.PersistentFlags().BoolP("link", "L", false, "Docker container linkage map")

	viper.BindPFlag("label", addCmd.PersistentFlags().Lookup("label"))
	viper.BindPFlag("env", addCmd.PersistentFlags().Lookup("env"))
	viper.BindPFlag("volume", addCmd.PersistentFlags().Lookup("volume"))
	viper.BindPFlag("port", addCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("link", addCmd.PersistentFlags().Lookup("link"))

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
	fmt.Println("labels: ", viper.Get("label"))
	fmt.Println("verbose: ", viper.Get("verbose"))
}
