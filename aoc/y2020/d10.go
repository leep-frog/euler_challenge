package y2020

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
)

func Day10() aoc.Day {
	return &day10{}
}

type day10 struct{}

func (d *day10) Solve(lines []string, o command.Output) {
}

func (d *day10) Cases() []*aoc.Case {
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
