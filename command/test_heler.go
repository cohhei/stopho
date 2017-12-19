package command

import (
	"bytes"

	"github.com/codegangsta/cli"
)

const Name string = "stopho"
const Version string = "0.2.0"

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
		Action:    CmdList,
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
		Action:    CmdVersion,
		Flags:     []cli.Flag{},
	},
	{
		Name:      "search",
		ShortName: "s",
		Usage:     "Generate search URLs from stock photos web sites",
		Action:    CmdSearch,
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
}

func InitAppForTests() (app *cli.App, outStream *bytes.Buffer) {
	app = cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Flags = GlobalFlags
	app.Commands = Commands
	outStream = new(bytes.Buffer)
	app.Writer = outStream

	return app, outStream
}
