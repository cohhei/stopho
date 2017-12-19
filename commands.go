package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kohei-kimura/stopho/command"
)

var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "japan, jp, j",
		Usage: "Enable Japan mode",
	},
}

var Commands = []cli.Command{
	{
		Name:      "list",
		ShortName: "l",
		Usage:     "List available stock photo web sites",
		Action:    command.CmdList,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "japan, j",
				Usage: "Enable Japan mode",
			},
			cli.BoolFlag{
				Name:  "quiet, q",
				Usage: "Only display URLs of stock photo sites",
			},
		},
	},
	{
		Name:      "version",
		ShortName: "v",
		Usage:     "Print the stopho version",
		Action:    command.CmdVersion,
		Flags:     []cli.Flag{},
	},
	{
		Name:      "search",
		ShortName: "s",
		Usage:     "Generate search URLs from stock photos web sites",
		Action:    command.CmdSearch,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "japan, j",
				Usage: "Enable Japan mode",
			},
			cli.BoolFlag{
				Name:  "quiet, q",
				Usage: "Display only URLs of stock photo sites",
			},
			cli.StringFlag{
				Name:  "site, s",
				Usage: "Display only sites matched `value`",
			},
		},
	},
	{
		Name:      "open",
		ShortName: "o",
		Usage:     "Open search URLs from stock photos web sites",
		Action:    command.CmdOpen,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "japan, j",
				Usage: "Enable Japan mode",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
