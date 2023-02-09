package y2016

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
)

func Day15() aoc.Day {
	return &day15{}
}

type day15 struct{}

func (d *day15) Solve(lines []string, o command.Output) {
}

func (d *day15) Cases() []*aoc.Case {
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
