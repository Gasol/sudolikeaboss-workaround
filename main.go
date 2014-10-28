package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "sudolikeaboss-server"
	app.Version = "0.1.0"
	app.Usage = "run the sudolikeaboss server for 1password5 workaround"
	app.Action = func(c *cli.Context) {
		runServer()
	}

	app.Run(os.Args)
}
