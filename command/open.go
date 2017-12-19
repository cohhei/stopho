package command

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/codegangsta/cli"
)

var openCommands = map[string]string{
	"darwin":  "open",
	"windows": "start",
	"linux":   "xdg-open",
}

func CmdOpen(c *cli.Context) {
	term := c.Args().Get(0)
	if term == "" {
		fmt.Fprintln(c.App.Writer, "Specify a search term!")
		cli.ShowCommandHelp(c, "open")
		return
	}

	sites := Sites
	if c.GlobalBool("japan") || c.Bool("japan") {
		sites = append(Sites, Sitesjp...)
	}

	openCommand := openCommands[runtime.GOOS]
	for _, site := range sites {
		url := fmt.Sprintf("%s%s", site.SearchUrl, term)
		fmt.Fprintf(c.App.Writer, "%s %s\n", openCommand, url)
		exec.Command(openCommand, url).Start()
	}
}
