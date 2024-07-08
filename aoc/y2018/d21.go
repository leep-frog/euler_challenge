package y2018

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
)

func Day21() aoc.Day {
	return &day21{}
}

type day21 struct{}

func (d *day21) Solve(lines []string, o command.Output) {
}

func (d *day21) Cases() []*aoc.Case {
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
