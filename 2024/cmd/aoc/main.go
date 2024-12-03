package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/keepri/adventofcode/internal/header"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Usage: aoc <day> <part>")
		os.Exit(0)
	}

	arg, err := strconv.ParseInt(args[0], 10, 0)
	if err != nil {
		fmt.Printf("Could not parse %v to int\n", args[0])
		os.Exit(0)
	}
	day := int(arg)

	arg, err = strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		fmt.Printf("Could not parse %v to int\n", args[1])
		os.Exit(0)
	}
	part := int(arg)

	header.Print(day, part)

	var answer int
	switch day {
	case 1:
		d := dayOne{}
		answer, err = d.solve(part)
	default:
		err = fmt.Errorf("Day %v not implemented\n", day)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Answer:", answer)
	}
}
