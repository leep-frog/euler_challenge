package y2016

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day20() aoc.Day {
	return &day20{}
}

type day20 struct{}

func (d *day20) Solve(lines []string, o command.Output) {
	var r *maths.Range
	for _, line := range lines {
		parts := strings.Split(line, "-")
		newR := maths.NewRange(parse.Atoi(parts[0]), parse.Atoi(parts[1]))
		if r == nil {
			r = newR
		} else {
			r = r.Merge(newR)
		}
	}

	pts := r.InflectionPoints()
	part1 := pts[1] + 1

	pts = append(pts, maths.Pow(2, 32))
	cnt := pts[0]
	for i := 2; i < len(pts); i += 2 {
		cnt += pts[i] - pts[i-1] - 1
	}

	o.Stdoutln(part1, cnt)
}

func (d *day20) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"3 4294967288",
			},
		},
		{
			ExpectedOutput: []string{
				"17348574 104",
			},
		},
	}
}
