package main

import (
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		print("Include a duration in seconds. Eg. 'gn 20' \n")
		return
	}

	argsWithoutProg := os.Args[1:]
	duration, _ := strconv.Atoi(argsWithoutProg[0])

	zzz(duration)
}
