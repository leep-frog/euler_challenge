package y2016

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/point"
	"github.com/leep-frog/euler_challenge/walker"
)

func Day01() aoc.Day {
	return &day01{}
}

type day01 struct{}

func (d *day01) Solve(lines []string, o command.Output) {
	path := strings.Split(lines[0], ", ")
	w := walker.CardinalWalker(walker.North, true)
	visited := map[string]bool{
		point.Origin[int]().String(): true,
	}
	var part2 int
	for _, dir := range path {
		if dir[0] == 'L' {
			w.Left()
		} else {
			w.Right()
		}

		for i := 0; i < parse.Atoi(dir[1:]); i++ {
			w.Walk(1)
			c := w.Position().String()
			if part2 == 0 && visited[c] {
				part2 = w.Position().ManhattanDistance(point.Origin[int]())
			}
			visited[c] = true
		}

	}
	o.Stdoutln(w.Position().ManhattanDistance(point.Origin[int]()), part2)
}

func (d *day01) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"8 4",
			},
		},
		{
			ExpectedOutput: []string{
				"301 130",
			},
		},
	}
}
