package y2017

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/walker"
)

func Day19() aoc.Day {
	return &day19{}
}

type day19 struct{}

func (d *day19) Solve(lines []string, o command.Output) {
	grid := parse.Split(lines, "")

	w := walker.CardinalWalker(walker.South, true)
	for i := 0; i < len(grid[0])-1; i++ {
		w.Move(walker.East, 1)
	}

	nextPointValid := func() (bool, string) {
		nextPoint := w.Position().Plus(w.CurrentVector())
		if nextPoint.Y >= 0 && nextPoint.X >= 0 && nextPoint.Y < len(grid) && nextPoint.X < len(grid[nextPoint.Y]) {
			return true, grid[nextPoint.Y][nextPoint.X]
		}
		return false, ""
	}

	var r []string
	for curChar, steps := "|", 1; ; steps++ {
		// Change direction
		if curChar == "+" {
			w.Left()
			if ok, c := nextPointValid(); !ok || c == " " {
				w.Right()
				w.Right()
			}
		}

		// Walk
		w.Walk(1)
		pos := w.Position()
		curChar = grid[pos.Y][pos.X]

		// Character processing
		if curChar != "-" && curChar != "|" && curChar != "+" && curChar != " " {
			r = append(r, curChar)
		}
		if curChar == " " {
			o.Stdoutln(strings.Join(r, ""), steps)
			return
		}
	}
}

func (d *day19) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"ABCDEF 38",
			},
		},
		{
			ExpectedOutput: []string{
				"LXWCKGRAOY 17302",
			},
		},
	}
}
