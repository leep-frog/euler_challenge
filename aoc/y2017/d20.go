package y2017

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/rgx"
)

func Day20() aoc.Day {
	return &day20{}
}

type day20 struct{}

func (d *day20) Solve(lines []string, o command.Output) {
	r := rgx.New(`^p=< *([-0-9]+), *([-0-9]+), *([-0-9]+)>, v=< *([-0-9]+), *([-0-9]+), *([-0-9]+)>, a=< *([-0-9]+), *([-0-9]+), *([-0-9]+)>$`)
	best := maths.Smallest[int, int]()
	for i, line := range lines {
		m := r.MustMatch(line)
		ax, ay, az := parse.Atoi(m[6]), parse.Atoi(m[7]), parse.Atoi(m[8])
		best.IndexCheck(i, maths.Abs(ax)+maths.Abs(ay)+maths.Abs(az))
	}
	o.Stdoutln(best.BestIndex())
}

func (d *day20) Cases() []*aoc.Case {
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
