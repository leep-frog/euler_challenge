package y2016

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
)

func Day17() aoc.Day {
	return &day17{}
}

type day17 struct{}

func (d *day17) Solve(lines []string, o command.Output) {
}

func (d *day17) Cases() []*aoc.Case {
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
