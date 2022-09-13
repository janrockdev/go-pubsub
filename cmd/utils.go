package cmd

import (
	"fmt"
	"github.com/janrockdev/go-pubsub/types"
	"github.com/janrockdev/go-pubsub/utils"
	"os"
)

func loadConfig() types.Config {
	path := loadConfigPath()
	config, err := types.ImportConfig(path)
	if err != nil {
		panic(err)
	}
	return config
}

func loadConfigPath() string {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		currentPath, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = currentPath + "/config.json"
		if utils.FileExists(path) == false {
			fmt.Printf("doesn't set config\n")
			os.Exit(1)
		}
	}
	return path
}
