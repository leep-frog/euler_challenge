package y2020

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day13() aoc.Day {
	return &day13{}
}

type day13 struct{}

func (d *day13) Solve(lines []string, o command.Output) {
	// Part 1 data
	best := maths.Smallest[int, int]()
	start := parse.Atoi(lines[0])

	// Part 2 data
	lp := maths.NewLinearProgression(0, 1)
	for mod, v := range strings.Split(lines[1], ",") {
		if v == "x" {
			continue
		}
		k := parse.Atoi(v)

		// Part 1 logic
		best.IndexCheck(k, (k - (start % k)))

		// Part 2 logic
		lp = lp.Merge(maths.NewLinearProgression(k-mod-1, k))
	}
	o.Stdoutln(best.BestIndex()*best.Best(), lp.Start().Plus(maths.One()))
}

func (d *day13) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"295 1068781",
			},
		},
		{
			ExpectedOutput: []string{
				"4135 640856202464541",
			},
		},
	}
}
