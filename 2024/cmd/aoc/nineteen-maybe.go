package main

import (
	"fmt"
	"strings"

	"github.com/keepri/adventofcode/internal/data"
)

const SEP = " => "

type dayNineteenMaybe struct{}

type tuple struct {
	From, To string
}

func (d dayNineteenMaybe) solve(part int) (int, error) {
	lines, err := data.GetLines(19)
	if err != nil {
		return 0, err
	}

	out := 0
	mappings := []tuple{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		onMapping := strings.Contains(line, SEP)
		if onMapping {
			mapping := strings.Split(line, SEP)
			from := mapping[0]
			to := mapping[1]
			tuple := tuple{from, to}
			mappings = append(mappings, tuple)
			continue
		}

		switch part {
		case 1:
			molecules := []string{}
			for _, m := range mappings {
				replacements := strings.Count(line, m.From)

				for nth := 1; nth <= replacements; nth++ {
					index := nthIndex(&line, m.From, nth)
					pre := line[:index]
					post := line[index+len(m.From):]
					molecule := pre + m.To + post

					if !exists(&molecules, molecule) {
						molecules = append(molecules, molecule)
					}
				}
			}
			out = len(molecules)
		case 2:
			return 0, fmt.Errorf("Day 19 part 2 not implemented... yet")
		default:
			return 0, fmt.Errorf("Invalid part: %v used for day 19", part)
		}
	}
	return out, nil
}
