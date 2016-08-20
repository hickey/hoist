package main

import (
	"fmt"
	"github.com/spf13/viper"
	"hoist/cmd"
	"hoist/store"
	"os"
)

func main() {
	// Initialize appl defaults
	initializeViper()

	store.Start()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	store.Stop()
}

func initializeViper() {
	viper.SetConfigName("hoist")
	viper.SetConfigType("TOML")
	viper.AddConfigPath("/etc/hoist")
	viper.AddConfigPath("$HOME/.hoist")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("hoist")
	viper.AutomaticEnv()

	viper.SetDefault("debug", false)
	viper.SetDefault("verbose", false)

}
