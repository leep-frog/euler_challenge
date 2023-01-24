package y2015

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
)

func Day01() aoc.Day {
	return &day01{}
}

type day01 struct{}

func (d *day01) Solve(lines []string, o command.Output) {
	floor := 0
	var basement int
	for i, c := range lines[0] {
		if c == '(' {
			floor++
		} else {
			floor--
		}
		if floor == -1 && basement == 0 {
			basement = i + 1
		}
	}
	o.Stdoutln(floor, basement)
}

func (d *day01) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"3 1",
			},
		},
		{
			ExpectedOutput: []string{
				"138 1771",
			},
		},
	}
}
