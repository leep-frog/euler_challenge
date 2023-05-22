package y2016

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/rgx"
)

func Day15() aoc.Day {
	return &day15{}
}

type day15 struct{}

type disc struct {
	numPositions int
	position     int
}

func (d *disc) String() string {
	return fmt.Sprintf("{%d,%d}", d.numPositions, d.position)
}

func (d *day15) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines, false), d.solve(lines, true))
}

func (d *day15) solve(lines []string, part2 bool) int {
	r := rgx.New("^Disc .* has ([1-9][0-9]*) positions; at time=([0-9]*), it is at position ([0-9]*).$")

	var discs []*disc
	for i, line := range lines {
		m := r.MustMatch(line)

		np := parse.Atoi(m[0])
		curp := (parse.Atoi(m[2]) + i) % np
		discs = append(discs, &disc{np, curp})
	}

	if part2 {
		discs = append(discs, &disc{11, len(lines) % 11})
	}

	lp := maths.NewLinearProgression(0, 1)
	for _, disc := range discs {
		lp = lp.Merge(maths.NewLinearProgression((disc.numPositions-disc.position)%disc.numPositions, disc.numPositions))
	}
	return lp.Start().ToInt() - 1
}

func (d *day15) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"5 85",
			},
		},
		{
			ExpectedOutput: []string{
				"121834 3208099",
			},
		},
	}
}
