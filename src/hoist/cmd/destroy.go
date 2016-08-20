package cmd

import (
	"github.com/spf13/cobra"
	"hoist/store"
)

func init() {
	RootCmd.AddCommand(destroyCmd)
}

var destroyCmd = &cobra.Command{
	Use:   "destroy NAME",
	Short: "Remove a Docker container being managed.",
	Run: func(cmd *cobra.Command, args []string) {
		destroyCommand(args)
	},
}

func destroyCommand(name []string) {
	store.DestroyStartup(name[0])
}
