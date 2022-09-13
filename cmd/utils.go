package cmd

import (
	"fmt"
	"github.com/janrockdev/go-pubsub/types"
	"github.com/janrockdev/go-pubsub/utils"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
)

func loadStringOrFilePath(c *cli.Context, inputFlagName string, inputFilePathFlagName string) []byte {
	var out []byte
	var err error
	inputString := c.String(inputFlagName)
	inputFilePath := c.String(inputFilePathFlagName)
	if inputString == "" {
		if inputFilePath == "" {
			fmt.Printf("need provide %s or %s\n", inputFlagName, inputFilePathFlagName)
			os.Exit(1)
		}
		out, err = ioutil.ReadFile(inputFilePath)
		if err != nil {
			fmt.Printf("read file occured error\n")
			os.Exit(1)
		}
	} else {
		out = []byte(inputString)
	}
	return out
}

func loadConfig() types.Config {
	path := loadConfigPath()
	config, err := types.ImportConfig(path)
	if err != nil {
		panic(err)
	}
	return config
}

func loadConfigPath() string {
	path := os.Getenv("ETHEREUM_WALLET_CONFIG_PATH")
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
