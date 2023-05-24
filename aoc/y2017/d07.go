package y2017

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
)

func Day07() aoc.Day {
	return &day07{}
}

type day07 struct{}

func (d *day07) Solve(lines []string, o command.Output) {
}

func (d *day07) Cases() []*aoc.Case {
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
