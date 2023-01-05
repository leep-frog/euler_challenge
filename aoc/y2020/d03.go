package y2020

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day03() aoc.Day {
	return &day03{}
}

type day03 struct{}

func (d *day03) slopeCount(grid [][]bool, xOffset, iOffset int) int {
	var sum int
	for i, x := 0, 0; i < len(grid); i, x = i+iOffset, x+xOffset {
		if grid[i][x%len(grid[i])] {
			sum++
		}
	}
	return sum
}

func (d *day03) Solve(lines []string, o command.Output) {
	grid := parse.AOCGrid(lines, false, true)

	parts := []int{
		d.slopeCount(grid, 1, 1),
		d.slopeCount(grid, 3, 1),
		d.slopeCount(grid, 5, 1),
		d.slopeCount(grid, 7, 1),
		d.slopeCount(grid, 1, 2),
	}
	o.Stdoutln(parts[1], maths.ProdSys(parts))
}

func (d *day03) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"7 336",
			},
		},
		{
			ExpectedOutput: []string{
				"203 3316272960",
			},
		},
	}
}
