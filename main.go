package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/pgrunde/invictus/create"
)

func main() {
	app := cli.NewApp()
	app.Name = "invictus"
	app.Usage = "generate templates"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Create a template.",
			Action: func(c *cli.Context) {
				fmt.Println("getting arg", c.Args().First())
				err := create.Project(c.Args().First())
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
