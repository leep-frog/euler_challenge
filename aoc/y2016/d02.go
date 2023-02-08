package y2016

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/point"
)

func Day02() aoc.Day {
	return &day02{}
}

type day02 struct{}

var (
	dirs = map[rune]*point.Point[int]{
		'R': point.New(0, 1),
		'D': point.New(1, 0),
		'L': point.New(0, -1),
		'U': point.New(-1, 0),
	}
)

func (d *day02) Solve(lines []string, o command.Output) {
	grids := [][][]string{
		{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		},
		{
			{"", "", "1", "", ""},
			{"", "2", "3", "4", ""},
			{"5", "6", "7", "8", "9"},
			{"", "A", "B", "C", ""},
			{"", "", "D", "", ""},
		},
	}
	starts := []*point.Point[int]{
		point.New(1, 1),
		point.New(2, 0),
	}

	var res []string
	for i := 0; i < 2; i++ {
		grid := grids[i]
		p := starts[i]

		var r []string
		for _, line := range lines {
			for _, c := range line {
				np := p.Plus(dirs[c])
				if np.X >= 0 && np.X < len(grid) && np.Y >= 0 && np.Y < len(grid) && grid[np.X][np.Y] != "" {
					p = np
				}
			}
			r = append(r, grid[p.X][p.Y])
		}
		res = append(res, strings.Join(r, ""))
	}
	o.Stdoutln(strings.Join(res, " "))
}

func (d *day02) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"1985 5DB3",
			},
		},
		{
			ExpectedOutput: []string{
				"74921 A6B35",
			},
		},
	}
}
