package y2020

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
)

func Day18() aoc.Day {
	return &day18{}
}

type day18 struct{}

func (d *day18) Solve(lines []string, o command.Output) {
}

func (d *day18) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"",
			},
		},
		{
			ExpectedOutput: []string{
				"",
			},
		},
	}
}
