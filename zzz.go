package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

func zzzCommand() *cli.Command {
	return &cli.Command{
		Name:    "zzz",
		Aliases: []string{"sleep"},
		Action: func(c *cli.Context) error {
			numMinutes, _ := strconv.Atoi(c.Args().First())
			// time.Sleep(time.Duration(numMinutes) * time.Minute)
			// fmt.Print(time.Duration(numMinutes) * time.Minute)
			// cmd := exec.Command("shutdown", "/h")
			// if err := cmd.Run(); err != nil {
			// 	panic(err)
			// }
			fmt.Printf("feeling sleepy... %d more minutes", numMinutes)
			return nil
		},
	}
}
