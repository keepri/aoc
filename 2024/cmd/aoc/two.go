package main

import (
	"fmt"
	"strings"

	"github.com/keepri/adventofcode/internal/data"
)

type dayTwo struct{}

const MAX_DIFF = 3

func (d dayTwo) solve(part int) (int, error) {
	lines, err := data.GetLines(2)
	if err != nil {
		return 0, err
	}

	safe := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		numbers := parseInts(split)

		switch part {
		case 1:
			if isSafe := d.safe(numbers); isSafe {
				safe++
			}
			return safe, nil
		case 2:
			perms := d.permutations(numbers)
			ackshually := false
			for perm := range perms {
				if isSafe := d.safe(perm); isSafe {
					ackshually = true
					break
				}
			}
			if ackshually {
				safe++
			}
		default:
			return 0, fmt.Errorf("Invalid part: %v used for day 1", part)
		}
	}

	return safe, nil
}

func (_ dayTwo) permutations(numbers []int) <-chan []int {
	ch := make(chan []int, len(numbers))
	go func() {
		defer close(ch)

		if len(numbers) < 3 {
			ch <- numbers
			return
		}

		for i := 0; i < len(numbers); i++ {
			perm := []int{}
			for j, n := range numbers {
				if i == j {
					continue
				}
				perm = append(perm, n)
			}
			ch <- perm
		}
	}()
	return ch
}

func (d dayTwo) safe(numbers []int) bool {
	out := true
	ascending := true
	for i := 0; i < len(numbers); i++ {
		if i == len(numbers)-1 {
			break
		}

		x := numbers[i]
		y := numbers[i+1]

		if i == 0 && x > y {
			ascending = false
		}

		if x == y ||
			(ascending && !d.validate(y-x)) ||
			(!ascending && !d.validate(x-y)) {
			out = false
			break
		}
	}
	return out
}

func (d dayTwo) validate(diff int) bool {
	if diff <= 0 || diff > MAX_DIFF {
		return false
	}
	return true
}
