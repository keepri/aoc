package main

import (
	"fmt"
	"os"

	"github.com/keepri/adventofcode/internal/header"
)

type aoc struct{}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("usage: ./aoc <day-number>")
		os.Exit(1)
	}

	aoc := aoc{}
	day := args[0]

	header.Print(day)

	switch day {
	case "1":
		aoc.day1()
	default:
		fmt.Println("day not done... yet")
	}
}
