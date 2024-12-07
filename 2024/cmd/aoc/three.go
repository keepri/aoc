package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/keepri/adventofcode/internal/data"
)

const (
	MUL         = "mul"
	PARAN_OPEN  = '('
	PARAN_CLOSE = ")"
	COMMA       = ","
)

type dayThree struct{}

func (d dayThree) solve(part int) (int, error) {
	lines, err := data.GetLines(3)
	if err != nil {
		return 0, err
	}

	out := 0
	for _, line := range lines {
		switch part {
		// TODO recurse
		case 1:
			split := strings.SplitAfter(line, MUL)
		chunk:
			for _, chunk := range split {
				if len(chunk) < 5 {
					continue
				}
				if chunk[0] != PARAN_OPEN {
					continue
				}
				params := strings.Split(chunk, COMMA)
				if len(params) < 2 {
					continue
				}
				firstParam := -1
				for i, param := range params {
					if len(param) < 2 {
						continue chunk
					}
					if i == 2 {
						continue chunk
					}
					if i == 0 {
						firstParam, err = strconv.Atoi(param[1:])
						if err != nil {
							continue chunk
						}
					} else {
						index := strings.Index(param, ")")
						if index == -1 {
							continue chunk
						}
						secondParam, err := strconv.Atoi(param[:index])
						if err != nil {
							continue chunk
						}
						out += firstParam * secondParam
					}
				}
			}
		default:
			return 0, fmt.Errorf("Invalid part: %v used for day 1", part)
		}
	}

	return out, nil
}

func (d dayThree) partOne(line string) (int, error) {
	return 0, nil
}
