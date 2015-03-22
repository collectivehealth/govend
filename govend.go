package main

import (
	"os"
	"runtime"

	"github.com/codegangsta/cli"
)

// USAGE
//
// COMMANDS: govend (maybe vend?), scan, prune, update, remove
//
// govend [global options..]
//        -v, -verbose
//

func main() {

	// do the max prox runtime CPU thang.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// start a new cli application.
	app := cli.NewApp()

	// define the cli application metadata.
	app.Name = "govend"
	app.Usage = "A CLI tool for vendoring golang packages."
	app.Version = "0.0.1"
	app.Author = "Jack Spirou"
	app.Email = "jack.spirou@me.com"

	// define the list of global options.
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "print things as they happen",
		},
	}

	// define the list of commands.
	app.Commands = []cli.Command{
		{
			Name:        "scan",
			Usage:       "Scans a go project for external package dependencies",
			ShortName:   "s",
			Description: "Use this command to find external package dependencies.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "write, w",
					Usage: "write the results to disk",
				},
				cli.StringFlag{
					Name:  "fmt, f",
					Usage: "format the results with values json, yaml, yml, or xml",
				},
			},
			Action: func(c *cli.Context) {
				if err := scancmd(c.Args().First(), c.String("write"), c.String("fmt")); err != nil {
					panic(err)
				}
			},
		},
	}

	// define the default action.
	app.Action = func(c *cli.Context) {
		if err := vend(c.GlobalBool("verbose")); err != nil {
			panic(err)
		}
	}

	// execute the cli command given.
	app.Run(os.Args)
}
