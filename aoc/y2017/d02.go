package y2017

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day02() aoc.Day {
	return &day02{}
}

type day02 struct{}

func (d *day02) Solve(lines []string, o command.Output) {
	grid := functional.Map(parse.Split(lines, "\t"), func(s []string) []int {
		return functional.Map(s, parse.Atoi)
	})
	var sum, sum2 int
	for _, row := range grid {
		sum += maths.Max(row...) - maths.Min(row...)
		for i, v := range row {
			for j, w := range row {
				if i == j {
					continue
				}
				if v > w {
					continue
				}
				if w%v == 0 {
					sum2 += w / v
				}
			}
		}
	}
	o.Stdoutln(sum, sum2)
}

func (d *day02) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"18 9",
			},
		},
		{
			ExpectedOutput: []string{
				"50376 267",
			},
		},
	}
}
