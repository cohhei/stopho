package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

func CmdList(c *cli.Context) {
	sites := Sites
	if c.GlobalBool("japan") || c.Bool("japan") {
		sites = append(Sites, Sitesjp...)
	}

	if !c.Bool("quiet") {
		table := tablewriter.NewWriter(c.App.Writer)
		table.SetHeader([]string{"Name", "URL"})
		for _, site := range sites {
			table.Append([]string{site.Name, site.Url})
		}
		table.Render()
		return
	}

	for _, site := range sites {
		fmt.Fprintf(c.App.Writer, "%s\n", site.Url)
	}
}
