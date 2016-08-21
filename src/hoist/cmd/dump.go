package cmd

import (
	"github.com/spf13/cobra"
	"hoist/store"
)

func init() {
	RootCmd.AddCommand(dumpCmd)
}

var dumpCmd = &cobra.Command{
	Use:    "dump NAME",
	Short:  "Dump the BoltDB contents for NAME.",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		dumpCommand(args)
	},
}

func dumpCommand(name []string) {
	store.DumpStartup(name[0])
}
