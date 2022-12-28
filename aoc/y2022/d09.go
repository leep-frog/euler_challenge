package y2022

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/point"
)

func Day09() aoc.Day {
	return &day09{}
}

type day09 struct{}

func (d *day09) solve(lines []string, o command.Output, ropeLength int) int {
	rope := []*point.Point[int]{}
	for i := 0; i < ropeLength; i++ {
		rope = append(rope, point.Origin[int]())
	}

	tail := func() *point.Point[int] {
		return rope[len(rope)-1]
	}

	m := map[string]bool{
		tail().String(): true,
	}

	dirMap := map[string]*point.Point[int]{
		"R": point.New(1, 0),
		"L": point.New(-1, 0),
		"U": point.New(0, 1),
		"D": point.New(0, -1),
	}

	diags := []*point.Point[int]{
		point.New(1, 1),
		point.New(1, -1),
		point.New(-1, 1),
		point.New(-1, -1),
	}

	var moves []*point.Point[int]
	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction, dist := dirMap[parts[0]], parse.Atoi(parts[1])
		for i := 0; i < dist; i++ {
			moves = append(moves, direction)
		}
	}

	for _, move := range moves {
		// Move the head
		rope[0] = rope[0].Plus(move)

		for i := 0; i < len(rope)-1; i++ {
			h := rope[i]
			t := rope[i+1]
			if h.ManhattanDistanceWithDiagonals(t) > 1 {
				if h.X == t.X {
					t.Y = (h.Y + t.Y) / 2
				} else if h.Y == t.Y {
					t.X = (h.X + t.X) / 2
				} else {
					for _, diag := range diags {
						move := diag.Plus(t)
						if move.ManhattanDistanceWithDiagonals(h) < 2 {
							rope[i+1] = move
							break
						}
					}
				}
			}
		}
		m[rope[len(rope)-1].String()] = true
	}
	return len(m)
}

func (d *day09) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines, o, 2), d.solve(lines, o, 10))
}

func (d *day09) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"13 1",
			},
		},
		{
			FileSuffix: "example2",
			ExpectedOutput: []string{
				"88 36",
			},
		},
		{
			ExpectedOutput: []string{
				"6098 2597",
			},
		},
	}
}
