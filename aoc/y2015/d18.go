package y2015

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day18() aoc.Day {
	return &day18{}
}

type day18 struct{}

func (d *day18) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines, false), d.solve(lines, true))
}

func (d *day18) solve(lines []string, part2 bool) int {
	grid := parse.AOCGrid(lines, false, true)

	if part2 {
		grid[0][0] = true
		grid[len(grid)-1][0] = true
		grid[0][len(grid[0])-1] = true
		grid[len(grid)-1][len(grid[0])-1] = true
	}

	for i := 0; i < 100; i++ {
		var newGrid [][]bool
		for i, row := range grid {
			var newRow []bool
			for j, c := range row {
				nc := parse.NeighborCount(grid, i, j, true)
				if c && nc < 2 || nc > 3 {
					c = false
				} else if !c && nc == 3 {
					c = true
				}
				newRow = append(newRow, c)
			}
			newGrid = append(newGrid, newRow)
		}
		grid = newGrid
		if part2 {
			grid[0][0] = true
			grid[len(grid)-1][0] = true
			grid[0][len(grid[0])-1] = true
			grid[len(grid)-1][len(grid[0])-1] = true
		}
	}
	return functional.Count2D(grid, true)
}

func (d *day18) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"4 7",
			},
		},
		{
			ExpectedOutput: []string{
				"821 886",
			},
		},
	}
}
