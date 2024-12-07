package main

import (
	"fmt"

	"github.com/keepri/adventofcode/internal/data"
)

type dayThree struct{}

func (d dayThree) solve(part int) (int, error) {
	lines, err := data.GetLines(3)
	if err != nil {
		return 0, err
	}

	safe := 0
	for _, line := range lines {
		switch part {
		case 1:
			fmt.Println(line)
		default:
			return 0, fmt.Errorf("Invalid part: %v used for day 1", part)
		}
	}

	return safe, nil
}
