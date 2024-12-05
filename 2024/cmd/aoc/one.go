package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/keepri/adventofcode/internal/data"
)

type dayOne struct{}

func (d dayOne) solve(part int) (int, error) {
	lines, err := data.GetLines(1)
	if err != nil {
		return 0, err
	}
	left, right := d.parse(&lines)

	switch part {
	case 1:
		sort.Ints(left)
		sort.Ints(right)

		out := []int{}
		for i := 0; i < len(left); i++ {
			leftLocId := left[i]
			rightLocId := right[i]

			if leftLocId >= rightLocId {
				dif := leftLocId - rightLocId
				out = append(out, dif)
			} else {
				dif := rightLocId - leftLocId
				out = append(out, dif)
			}
		}

		sum := 0
		for _, v := range out {
			sum += v
		}

		return sum, nil
	case 2:
		similarity := 0
		for i := 0; i < len(left); i++ {
			locId := left[i]
			occurences := d.occurences(&right, &locId)
			similarity += locId * occurences
		}
		return similarity, nil
	default:
		return 0, fmt.Errorf("Invalid part: %v used for day 1", part)
	}
}

func (_ *dayOne) parse(lines *[]string) ([]int, []int) {
	left := []int{}
	right := []int{}
	for _, line := range *lines {
		v := strings.Split(line, "   ")
		ints := parseInts(v)

		left = append(left, ints[0])
		right = append(right, ints[1])
	}
	return left, right
}

func (_ *dayOne) occurences(arr *[]int, locId *int) int {
	count := 0
	for _, n := range *arr {
		if n == *locId {
			count++
		}
	}
	return count
}
