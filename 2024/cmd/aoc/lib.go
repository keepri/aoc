package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseInts(args []string) [2]int {
	if len(args) < 2 {
		fmt.Println("Usage: aoc <day> <part>")
		os.Exit(0)
	}

	out := [2]int{}
	for _, a := range args {
		arg, err := strconv.ParseInt(a, 10, 0)
		if err != nil {
			fmt.Printf("Could not parse %v to int\n", a)
			os.Exit(0)
		}
		out[0] = int(arg)
		out[1] = int(arg)
	}

	return out
}
