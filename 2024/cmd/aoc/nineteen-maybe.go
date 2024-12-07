package main

import (
	"fmt"
	"strings"

	"github.com/keepri/adventofcode/internal/data"
)

const SEP = " => "
const ELECTRON = "e"

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
			steps := 0
			d.reverseEngineer(&line, &mappings, &steps)
			return steps, nil
		default:
			return 0, fmt.Errorf("Invalid part: %v used for day 19", part)
		}
	}
	return out, nil
}

func (d dayNineteenMaybe) reverseEngineer(molecule *string, mappings *[]tuple, steps *int) {
	if *molecule == ELECTRON {
		return
	}

	for _, m := range *mappings {
		if !strings.Contains(*molecule, m.To) {
			continue
		}

		mol := strings.Replace(*molecule, m.To, m.From, 1)
		if m.From == ELECTRON && mol != ELECTRON {
			continue
		}
		*steps++
		*molecule = mol
		if mol == ELECTRON {
			break
		}
	}

	d.reverseEngineer(molecule, mappings, steps)
}
