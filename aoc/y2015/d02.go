package y2015

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day02() aoc.Day {
	return &day02{}
}

type day02 struct{}

func (d *day02) Solve(lines []string, o command.Output) {
	var sum1, sum2 int
	for _, row := range parse.ToGrid(lines, "x") {
		l, w, h := row[0], row[1], row[2]
		sum1 += 2*(l*w+l*h+w*h) + maths.Min(l*w, l*h, w*h)
		sum2 += 2*maths.Min(l+w, l+h, w+h) + l*w*h
	}
	o.Stdoutln(sum1, sum2)
}

func (d *day02) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"101 48",
			},
		},
		{
			ExpectedOutput: []string{
				"1606483 3842356",
			},
		},
	}
}
