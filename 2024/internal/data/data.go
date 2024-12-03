package data

import (
	"bufio"
	"fmt"
	"os"
)

func GetLines(day int) ([]string, error) {
	var f *os.File
	switch day {
	case 1:
		if file, err := os.Open("internal/data/one"); err != nil {
			fmt.Printf("Error opening file: %v\n", err)
		} else {
			f = file
		}
	default:
		return nil, fmt.Errorf("Invalid day: %v passed into get line\n", day)
	}
	defer f.Close()

	out := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		out = append(out, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return out, nil
}
