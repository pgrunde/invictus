package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/pgrunde/invictus/create"
)

// start shit

func main() {
	var dbname string
	var dbuser string
	var dbpw string
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
		cli.StringFlag{
			Name:        "dbuser",
			Usage:       "set the database user when creating a new project",
			Destination: &dbuser,
		},
		cli.StringFlag{
			Name:        "dbpw",
			Usage:       "set the database password when creating a new project",
			Destination: &dbpw,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Generates a new http server",
			Action: func(c *cli.Context) {
				err := create.NewProject(c.Args().First(), dbname, dbuser, dbpw)
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
