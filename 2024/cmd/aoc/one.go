package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/keepri/adventofcode/internal/data"
)

type dayOne struct{}

func (d *dayOne) solve(part int) (int, error) {
	lines, err := data.GetLines(1)
	if err != nil {
		return 0, err
	}
	left, right := d.parse(&lines)

	switch part {
	case 1:
		sort.Ints(left)
		sort.Ints(right)

		result := []int{}
		for i := 0; i < len(left); i++ {
			leftLocId := left[i]
			rightLocId := right[i]

			if leftLocId >= rightLocId {
				dif := leftLocId - rightLocId
				result = append(result, dif)
			} else {
				dif := rightLocId - leftLocId
				result = append(result, dif)
			}
		}

		answer := 0
		for _, v := range result {
			answer += v
		}

		return answer, nil
	case 2:
		similarity := 0
		for i := 0; i < len(left); i++ {
			locId := left[i]
			occurences := d.occurences(&right, &locId)
			similarity += locId * occurences
		}
		return similarity, nil
	}

	return 0, fmt.Errorf("Invalid part: %v used for day 1", part)
}

func (_ *dayOne) parse(lines *[]string) ([]int, []int) {
	left := []int{}
	right := []int{}
	for _, line := range *lines {
		v := strings.Split(line, "   ")
		one := v[0]
		two := v[1]

		n, err := strconv.ParseInt(one, 10, 0)
		if err != nil {
			panic(err)
		}
		left = append(left, int(n))

		m, err := strconv.ParseInt(two, 10, 0)
		if err != nil {
			panic(err)
		}
		right = append(right, int(m))
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
