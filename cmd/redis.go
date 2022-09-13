package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var (
	RedisCmd = &cli.Command{
		Name:        "ping",
		Usage:       "",
		Description: "",
		//ArgsUsage:   "<raw> <rawfile>", //flags
		//Flags: []cli.Flag{
		//	rawFlag,
		//	rawFileFlag,
		//},
		Action: func(c *cli.Context) error {
			config := loadConfig()
			fmt.Println(config)
			return nil
		},
	}

	RedisCommand = &cli.Command{
		Name:        "redis",
		Usage:       "RedisDB commands",
		ArgsUsage:   "",
		Category:    "Commands",
		Description: "",
		Subcommands: []*cli.Command{
			RedisCmd,
		},
	}
)
