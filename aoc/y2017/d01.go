package y2017

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
)

func Day01() aoc.Day {
	return &day01{}
}

type day01 struct{}

func (d *day01) Solve(lines []string, o command.Output) {
}

func (d *day01) Cases() []*aoc.Case {
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
