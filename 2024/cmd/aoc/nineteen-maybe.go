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
			molecules := d.createMolecules(line, mappings)
			out = len(molecules)
		case 2:
			steps := d.reverseEngineer(line, &mappings, 0)
			out = steps
		default:
			return 0, fmt.Errorf("Invalid part: %v used for day 19", part)
		}
	}

	return out, nil
}

func (_ dayNineteenMaybe) createMolecules(molecule string, mappings []tuple) []string {
	molecules := []string{}
	for _, m := range mappings {
		replacements := strings.Count(molecule, m.From)
		for nth := 1; nth <= replacements; nth++ {
			index := nthIndex(molecule, m.From, nth)
			pre := molecule[:index]
			post := molecule[index+len(m.From):]
			molecule := pre + m.To + post
			if !exists(molecules, molecule) {
				molecules = append(molecules, molecule)
			}
		}
	}
	return molecules
}

func (d dayNineteenMaybe) reverseEngineer(
	molecule string,
	mappings *[]tuple,
	steps int,
) int {
	if molecule == ELECTRON {
		return steps
	}

	for _, m := range *mappings {
		if (m.From == ELECTRON && m.To != molecule) ||
			(!strings.Contains(molecule, m.To)) {
			continue
		}
		molecule = strings.Replace(molecule, m.To, m.From, 1)
		steps++
		if molecule == ELECTRON {
			break
		}
	}

	return d.reverseEngineer(molecule, mappings, steps)
}
