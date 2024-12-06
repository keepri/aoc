package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInts(args []string) []int {
	out := []int{}
	for _, a := range args {
		arg, err := strconv.ParseInt(a, 10, 0)
		if err != nil {
			fmt.Printf("Could not parse %v to int\n", a)
			os.Exit(0)
		}
		out = append(out, int(arg))
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
