package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var (
	RedisCmd = &cli.Command{
		Name:        "sendrawtx",
		Usage:       "send raw tranasaction from node",
		Description: "send raw tranasaction from node",
		ArgsUsage:   "<raw> <rawfile>",
		Flags: []cli.Flag{
			rawFlag,
			rawFileFlag,
		},
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
		Category:    "Utility Commands",
		Description: "",
		Subcommands: []*cli.Command{
			RedisCmd,
		},
	}
)
