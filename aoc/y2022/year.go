package y2022

import (
	"github.com/leep-frog/euler_challenge/aoc/aoc"
)

func Year() *aoc.Year {
	return &aoc.Year{
		Number: 2022,
		Days: []aoc.Day{
			Day01(),
			Day02(),
			Day03(),
			Day04(),
			Day05(),
			Day06(),
			Day07(),
			Day08(),
			Day09(),
			Day10(),
			Day11(),
			Day12(),
			Day13(),
			Day14(),
			Day15(),
			Day16(),
			Day17(),
			Day18(),
			Day19(),
			// END_OF_DAYS
		},
	}
}
