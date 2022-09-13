package cmd

import (
	"github.com/janrockdev/go-pubsub/service"
	"github.com/urfave/cli/v2"
)

var (
	// RedisCmd Commands
	RedisCmd = &cli.Command{
		Name:        "ping",
		Usage:       "redis ping <args>",
		Description: "",
		//ArgsUsage:   "<raw> <rawfile>", //flags
		//Flags: []cli.Flag{
		//	rawFlag,
		//	rawFileFlag,
		//},
		Action: func(c *cli.Context) error {
			config := loadConfig()
			service.RedisService(config.ServerUrl)
			return nil
		},
	}

	// RedisCommand Global Options
	RedisCommand = &cli.Command{
		Name:        "redis",
		Usage:       "RedisDB commands",
		ArgsUsage:   "",
		Category:    "Database Commands",
		Description: "Commands for interaction with RedisDB",
		Subcommands: []*cli.Command{
			RedisCmd,
		},
	}
)
