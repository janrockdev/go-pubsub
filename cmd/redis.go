package cmd

import (
	"github.com/janrockdev/go-pubsub/service"
	"github.com/janrockdev/go-pubsub/utils"
	"github.com/urfave/cli/v2"
)

var (
	// RedisCmd Command
	RedisCmd = &cli.Command{
		Name:        "ping",
		Usage:       "redis ping <args>",
		Description: "",
		ArgsUsage:   "<filter>",
		Flags: []cli.Flag{
			filter,
		},
		Action: func(c *cli.Context) error {
			config := loadConfig()
			err := service.RedisService(config.ServerUrl, c.String("filter"))
			if err != nil {
				utils.Logr.Error(err)
				return err
			}
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
