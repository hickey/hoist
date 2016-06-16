package main

import (
	"fmt"
	"hoist/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("this is a test")

}
