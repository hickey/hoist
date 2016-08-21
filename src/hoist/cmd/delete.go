package cmd

import (
	"github.com/spf13/cobra"
	"hoist/store"
)

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().BoolP("label", "l", false, "Delete a label")
	deleteCmd.PersistentFlags().BoolP("env", "e", false, "Delete environmental variable")
	deleteCmd.PersistentFlags().BoolP("volume", "V", false, "Delete volume mapping")
	deleteCmd.PersistentFlags().BoolP("port", "p", false, "Delete Docker port mapping")
	deleteCmd.PersistentFlags().BoolP("link", "L", false, "Delete Docker container linkage map")

	viper.BindPFlag("label", deleteCmd.PersistentFlags().Lookup("label"))
	viper.BindPFlag("env", deleteCmd.PersistentFlags().Lookup("env"))
	viper.BindPFlag("volume", deleteCmd.PersistentFlags().Lookup("volume"))
	viper.BindPFlag("port", deleteCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("link", deleteCmd.PersistentFlags().Lookup("link"))
}

var deleteCmd = &cobra.Command{
	Use:   "delete NAME",
	Short: "Delete settings from a Docker container.",
	Run: func(cmd *cobra.Command, args []string) {
		deleteCommand(args)
	},
}

func deleteCommand(name []string) {

}
