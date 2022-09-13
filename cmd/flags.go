package cmd

import "github.com/urfave/cli/v2"

var (
	valueFlag = &cli.StringFlag{
		Name:     "value",
		Usage:    "send value",
		Required: true,
	}
	portFlag = &cli.IntFlag{
		Name:  "port",
		Usage: "port",
		Value: 8080,
	}
)
