package main

import (
	"fmt"
	"strings"

	"github.com/keepri/adventofcode/internal/data"
)

const sep = " => "

type dayNineteenMaybe struct{}

func (d dayNineteenMaybe) solve(part int) (int, error) {
	lines, err := data.GetLines(19)
	if err != nil {
		return 0, err
	}

	out := 0
	mappings := make(map[string][]string, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		onMapping := strings.Contains(line, sep)
		if onMapping {
			mapping := strings.Split(line, sep)
			from := mapping[0]
			to := mapping[1]
			if _, ok := mappings[from]; ok {
				mappings[from] = append(mappings[from], to)
			} else {
				mappings[from] = []string{to}
			}
			continue
		}

		switch part {
		case 1:
			molecules := []string{}
			for from, to := range mappings {
				replacements := strings.Count(line, from)
				for _, t := range to {
					for nth := 1; nth <= replacements; nth++ {
						index := nthIndex(line, from, nth)
						pre := line[:index]
						post := line[index+len(from):]

						molecule := pre + t + post
						if exists(molecules, molecule) {
							continue
						}
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
