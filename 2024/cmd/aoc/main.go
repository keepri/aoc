package main

import (
	"fmt"
	"os"
	"time"

	"github.com/keepri/adventofcode/internal/header"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: aoc <day> <part>")
		os.Exit(0)
	}
	ints := parseInts(args)
	day, part := ints[0], ints[1]

	header.Print(day, part)

	start := time.Now()
	var err error
	var answer int
	switch day {
	case 1:
		d := dayOne{}
		answer, err = d.solve(part)
	case 2:
		d := dayTwo{}
		answer, err = d.solve(part)
	case 3:
		d := dayThree{}
		answer, err = d.solve(part)
	case 19:
		d := dayNineteenMaybe{}
		answer, err = d.solve(part)
	default:
		err = fmt.Errorf("Day %v not implemented\n", day)
	}
	elapsed := float64(time.Since(start).Abs().Microseconds()) / 1000

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Answer:", answer, "in", elapsed, "ms")
	}
}
