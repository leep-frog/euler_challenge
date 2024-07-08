package y2016

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
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

	var r1, r2 []string
	for _, m := range counts {
		best1 := maths.Largest[rune, int]()
		best2 := maths.Smallest[rune, int]()
		for k, v := range m {
			best1.IndexCheck(k, v)
			best2.IndexCheck(k, v)
		}
		r1 = append(r1, string(best1.BestIndex()))
		r2 = append(r2, string(best2.BestIndex()))
	}
	o.Stdoutln(strings.Join(r1, ""), strings.Join(r2, ""))
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
