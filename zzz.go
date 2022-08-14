package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

func pluralize(num int, value string) string {
	if num > 1 || num == 0 {
		return value + "s"
	} else {
		return value
	}
}

func gnCommand() *cli.Command {
	return &cli.Command{
		Name:    "gn",
		Aliases: []string{"goodnight"},
		Action: func(c *cli.Context) error {
			numMinutes, _ := strconv.Atoi(c.Args().First())
			fmt.Printf("( ु⁎ᴗ_ᴗ⁎)ु.｡oO %d %s", numMinutes, pluralize(numMinutes, "minute"))

			time.Sleep(time.Duration(numMinutes) * time.Minute)

			cmd := sleep()

			if err := cmd.Run(); err != nil {
				panic(err)
			}
			return nil
		},
	}
}
