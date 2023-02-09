package y2016

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/maths"
)

func Day06() aoc.Day {
	return &day06{}
}

type day06 struct{}

func (d *day06) Solve(lines []string, o command.Output) {
	var counts []map[rune]int
	for i := 0; i < len(lines[0]); i++ {
		counts = append(counts, map[rune]int{})
	}
	for _, line := range lines {
		for col, c := range line {
			counts[col][c]++
		}
	}

	var r []string
	for _, m := range counts {
		best := maths.Smallest[rune, int]()
		for k, v := range m {
			best.IndexCheck(k, v)
		}
		r = append(r, string(best.BestIndex()))
	}
	fmt.Println(strings.Join(r, ""))
}

func (d *day06) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"easter advent",
			},
		},
		{
			ExpectedOutput: []string{
				"zcreqgiv pljvorrk",
			},
		},
	}
}
