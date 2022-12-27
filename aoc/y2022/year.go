package y2022

import (
	"github.com/leep-frog/euler_challenge/aoc/aoc"
)

func Year() *aoc.Year {
	return &aoc.Year{
		Number: 2022,
		Days:   []aoc.Day{
			Day01(),
			// END_OF_DAYS
		},
	}
}
