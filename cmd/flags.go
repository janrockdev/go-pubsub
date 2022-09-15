package cmd

import "github.com/urfave/cli/v2"

var (
	filter = &cli.StringFlag{
		Name:     "filter",
		Aliases:  []string{"f"},
		Usage:    "filter",
		Required: true,
		Value:    "channel_*",
		// Destination: &f,
		// EnvVars: []string{"APP_FILTER"},
	}
)
