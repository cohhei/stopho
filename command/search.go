package command

import (
	"fmt"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

func CmdSearch(c *cli.Context) {
	term := c.Args().Get(0)
	if term == "" {
		fmt.Fprintln(c.App.Writer, "Specify a search term!")
		cli.ShowCommandHelp(c, "search")
		return
	}

	sites := Sites
	if c.GlobalBool("japan") || c.Bool("japan") {
		sites = append(Sites, Sitesjp...)
	}

	if c.String("site") != "" {
		edited := []Site{}
		for _, site := range sites {
			if strings.Contains(site.Name, c.String("site")) {
				edited = append(edited, site)
			}
		}

		if len(edited) <= 0 {
			fmt.Fprintf(c.App.Writer, "No site matched `%s`\n", c.String("site"))
			cli.ShowCommandHelp(c, "search")
			return
		}
		sites = edited
	}

	if !c.Bool("quiet") {
		table := tablewriter.NewWriter(c.App.Writer)
		table.SetHeader([]string{"Name", "URL"})

		for _, site := range sites {
			table.Append([]string{site.Name, site.SearchUrl + c.Args().Get(0)})
		}
		table.Render()
		return
	}

	for _, site := range sites {
		fmt.Fprintf(c.App.Writer, "%s%s\n", site.SearchUrl, c.Args().Get(0))
	}
}
