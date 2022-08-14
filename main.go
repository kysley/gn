package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:        "gn",
		Description: "goodnight, get some rest",
		Commands:    []*cli.Command{gnCommand()},
	}
	app.Run(os.Args)
}
