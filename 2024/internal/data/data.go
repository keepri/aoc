package data

import (
	"bufio"
	"fmt"
	"os"
)

func GetLines(day int) ([]string, error) {
	var f *os.File
	var err error
	switch day {
	case 1:
		f, err = openFile("internal/data/one")
	case 2:
		f, err = openFile("internal/data/two")
	case 19:
		f, err = openFile("internal/data/nineteen-maybe")
	default:
		return nil, fmt.Errorf("Invalid day: %v passed into get line\n", day)
	}
	defer f.Close()

	if err != nil {
		return []string{}, err
	}

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

func openFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
