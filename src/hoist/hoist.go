package main

import (
	"fmt"
	"hoist/cmd"
	"hoist/store"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	store.Start()
}
