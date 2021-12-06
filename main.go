package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:        "Zzz",
		Description: "I'm feeling sleepy...",
		Commands:    []*cli.Command{zzzCommand()},
	}
	app.Run(os.Args)
}
