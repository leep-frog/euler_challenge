package y2020

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
)

func Day20() aoc.Day {
	return &day20{}
}

type day20 struct{}

func (d *day20) Solve(lines []string, o command.Output) {
}

func (d *day20) Cases() []*aoc.Case {
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
