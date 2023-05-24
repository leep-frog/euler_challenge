package y2017

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
)

func Day09() aoc.Day {
	return &day09{}
}

type day09 struct{}

func (d *day09) Solve(lines []string, o command.Output) {
}

func (d *day09) Cases() []*aoc.Case {
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
