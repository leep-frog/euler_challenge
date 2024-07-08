package y2020

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/point"
	"github.com/leep-frog/euler_challenge/walker"
)

func Day12() aoc.Day {
	return &day12{}
}

type day12 struct{}

func (d *day12) Solve(lines []string, o command.Output) {
	dirMap := map[byte]walker.CardinalDirection{
		'N': walker.North,
		'E': walker.East,
		'S': walker.South,
		'W': walker.West,
	}

	// Part 1 elements
	w := walker.CardinalWalker(walker.East, false)

	// Part 2 elements
	wayPoint := point.New(10, 1)
	start2 := point.Origin[int]()

	for _, line := range lines {
		d, v := line[0], parse.Atoi(line[1:])
		switch d {
		case 'F':
			w.Walk(v)
			start2 = start2.Plus(wayPoint.Times(v))
		case 'R':
			for i := 0; i < v; i += 90 {
				w.Right()
				wayPoint = wayPoint.Rotate(true)
			}
		case 'L':
			for i := 0; i < v; i += 90 {
				w.Left()
				wayPoint = wayPoint.Rotate(false)
			}
		default:
			w.Move(dirMap[d], v)
			wayPoint = wayPoint.Plus(w.GetVector(dirMap[d]).Times(v))
		}
	}
	o.Stdoutln(
		w.Position().ManhattanDistance(point.Origin[int]()),
		start2.ManhattanDistance(point.Origin[int]()),
	)
}

func (d *day12) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"25 286",
			},
		},
		{
			ExpectedOutput: []string{
				"2458 145117",
			},
		},
	}
}
