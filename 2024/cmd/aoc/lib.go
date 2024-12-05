package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInts(args []string) [2]int {
	if len(args) < 2 {
		fmt.Println("Usage: aoc <day> <part>")
		os.Exit(0)
	}

	out := [2]int{}
	for i, a := range args {
		arg, err := strconv.ParseInt(a, 10, 0)
		if err != nil {
			fmt.Printf("Could not parse %v to int\n", a)
			os.Exit(0)
		}
		out[i] = int(arg)
	}

	return out
}

func nthIndex(str, substr string, n int) int {
	if n <= 0 {
		return -1
	}

	index := -1
	start := 0
	for i := 0; i < n; i++ {
		pos := strings.Index(str[start:], substr)
		if pos == -1 {
			return pos
		}

		index = start + pos
		start = index + 1
	}
	return index
}

func exists(slice []string, str string) bool {
	exists := false
	for _, v := range slice {
		if v == str {
			exists = true
			break
		}
	}
	return exists
}
