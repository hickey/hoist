package cmd

import (
	"app"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(createCmd)

}

var createCmd = &cobra.Command{
	Use:   "create NAME CONTAINER",
	Short: "Create a Docker container startup script",
	Run: func(cmd *cobra.Command, args []string) {
		createCommand(args)
	},
}

func createCommand(args []string) {
	name := args[0]
	cont := args[1]

	log := app.GetLogger()

	log.Warningf("name = %s", name)
	log.Warningf("cont = %s", cont)

}
