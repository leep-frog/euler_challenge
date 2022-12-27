package y2022

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
)

func Day01() aoc.Day {
	return &day01{}
}

type day01 struct{}

func (d *day01) Solve1(lines []string, o command.Output) {
}

func (d *day01) Solve2(lines []string, o command.Output) {
}

func (d *day01) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			ExpectedOutput: "",
		},
		{
			FileSuffix:     "example",
			ExpectedOutput: "",
		},
	}
}
