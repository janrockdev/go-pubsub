package cmd

import (
	"github.com/janrockdev/go-pubsub/service"
	"github.com/janrockdev/go-pubsub/utils"
	"github.com/urfave/cli/v2"
)

var (
	// NatsCmd Command
	NatsCmd = &cli.Command{
		Name:        "ping",
		Usage:       "nats ping <args>",
		Description: "",
		ArgsUsage:   "<filter>",
		Flags: []cli.Flag{
			filter,
		},
		Action: func(c *cli.Context) error {
			config := loadConfig()
			err := service.NatsService(config.ServerUrl, c.String("filter"))
			if err != nil {
				utils.Logr.Error(err)
				return err
			}
			return nil
		},
	}

	// NatsCommand Global Options
	NatsCommand = &cli.Command{
		Name:        "nats",
		Usage:       "NATS/STAN commands",
		ArgsUsage:   "",
		Category:    "Messaging System Commands",
		Description: "Commands for interaction with NATS/STAN",
		Subcommands: []*cli.Command{
			NatsCmd,
		},
	}
)
