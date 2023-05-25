package y2017

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/hexgrid"
	"github.com/leep-frog/euler_challenge/maths"
)

func Day11() aoc.Day {
	return &day11{}
}

type day11 struct{}

func (d *day11) Solve(lines []string, o command.Output) {
	tile := hexgrid.Origin()
	max := maths.Largest[int, int]()
	for _, path := range strings.Split(lines[0], ",") {
		tile.MoveCode(path)
		max.Check(tile.Distance())
	}
	o.Stdoutln(tile.Distance(), max.Best())
}

func (d *day11) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"",
			},
		},
		{
			ExpectedOutput: []string{
				"670 1426",
			},
		},
	}
}
