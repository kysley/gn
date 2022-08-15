package main

import (
	"fmt"
	"time"
)

func pluralize(num int, value string) string {
	if num > 1 || num == 0 {
		return value + "s"
	} else {
		return value
	}
}

func zzz(duration int) {
	fmt.Printf("( ु⁎ᴗ_ᴗ⁎)ु.｡oO %d %s", duration, pluralize(duration, "minute"))

	time.Sleep(time.Duration(duration) * time.Minute)

	cmd := sleep()

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
