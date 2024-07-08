package y2015

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day25() aoc.Day {
	return &day25{}
}

type day25 struct{}

func (d *day25) Solve(lines []string, o command.Output) {
	code := 20151125
	wantRow, wantCol := parse.Atoi(lines[0]), parse.Atoi(lines[1])
	for i, row, column := 1, 1, 1; ; i++ {
		if row == wantRow && column == wantCol {
			o.Stdoutln(code)
			return
		}
		if row == 1 {
			row, column = column+1, 1
		} else {
			column++
			row--
		}
		code = (code * 252533) % 33554393
	}
}

func (d *day25) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"31527494",
			},
		},
		{
			ExpectedOutput: []string{
				"2650453",
			},
		},
	}
}
