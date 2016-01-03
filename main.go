package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/pgrunde/invictus/create"
)

func main() {
	var dbname string
	app := cli.NewApp()
	app.Name = "invictus"
	app.Usage = "generate templates"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "dbname",
			Usage:       "set the database name when creating a new project",
			Destination: &dbname,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Create a template.",
			Action: func(c *cli.Context) {
				err := create.NewProject(c.Args().First(), dbname)
				if err != nil {
					fmt.Println(err)
				}
			},
		},
	}
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}
	app.Run(os.Args)
}
