package cmd

import (
	"github.com/janrockdev/go-pubsub/service"
	"github.com/urfave/cli/v2"
)

var (
	// NatsCmd Command
	NatsCmd = &cli.Command{
		Name:        "channels",
		Usage:       "list all nats channels <args>",
		Description: "",
		ArgsUsage:   "<filter>",
		Flags: []cli.Flag{
			filter,
		},
		Action: func(c *cli.Context) error {
			config := loadConfig()
			service.NatsService(config.ServerUrl, c.String("filter"))
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
