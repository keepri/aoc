package header

import "fmt"

type printer struct{}

func Print(day string) {
	fmt.Println("===========================")
	fmt.Println("| 🚀 AOC 2024 LETS GO! 🌲 |")
	fmt.Println("===========================")

	p := printer{}

	switch day {
	case "1":
		p.day1()
	}

	fmt.Println()
}
