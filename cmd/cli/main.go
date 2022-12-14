package main

import (
	"fmt"
	"github.com/janrockdev/go-pubsub/cmd"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "Go-PubSub Application",
		Usage: "(app description)",
		Commands: []*cli.Command{
			cmd.RedisCommand,
			cmd.NatsCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		_ = fmt.Errorf("%s", err)
	}
}
