package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"hoist/store"
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

func createCommand(args []string) error {
	if len(args) != 2 {
		log.Error("Docker container needs to be specified")
		return errors.New("bad arg")
	}
	name := args[0]
	cont := args[1]

	return store.CreateStartup(name, cont)
}
