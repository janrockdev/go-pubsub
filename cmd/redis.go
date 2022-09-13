package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	sendrawtxCmd = &cli.Command{
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
			wallet := wallet.ImportEmptyEthereumWallet(config)
			rawb := loadStringOrFilePath(c, "raw", "rawfile")
			res, err := wallet.SendRawTransaction(string(rawb))
			if err != nil {
				fmt.Printf("sendRawTransaction error: %s\n", err)
				os.Exit(1)
			}
			fmt.Println(res)
			return nil
		},
	}

	NodeCommand = &cli.Command{
		Name:      "node",
		Usage:     "Ethereum node commands",
		ArgsUsage: "",
		Category:  "Node Commands",
		Description: "you can use node command to connect node and get information about ethereum, if you want to connect node," +
			"you must set node_url and network in config.json. if you want to get transaction history(normal, internal, erc20) by an address," +
			"you also need set etherscan_api_key in config.json",
		Subcommands: []*cli.Command{
			sendrawtxCmd,
		},
	}
)
