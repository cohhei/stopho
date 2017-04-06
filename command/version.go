package command

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func CmdVersion(c *cli.Context) {
	fmt.Fprintf(c.App.Writer, "%s version %s\n", c.App.Name, c.App.Version)
}
