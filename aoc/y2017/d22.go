package y2017

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
)

func Day22() aoc.Day {
	return &day22{}
}

type day22 struct{}

func (d *day22) Solve(lines []string, o command.Output) {
}

func (d *day22) Cases() []*aoc.Case {
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
