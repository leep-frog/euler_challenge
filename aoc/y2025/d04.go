package y2025

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day04() aoc.Day {
	return &day04{}
}

type day04 struct{}

func (d *day04) Solve(lines []string, o command.Output) {
	grid := parse.MapToGrid(lines, map[rune]bool{'@': true, '.': false})

	var rolls int

	removed := true
	for removed {
		removed = false
		var toRemove [][2]int

		for x, row := range grid {
			for y, value := range row {
				if !value {
					continue
				}

				neighborCount := parse.NeighborCount(grid, x, y, true)
				if neighborCount < 4 {
					rolls++
					toRemove = append(toRemove, [2]int{x, y})
					removed = true
				}

			}
		}

		for _, coord := range toRemove {
			grid[coord[0]][coord[1]] = false
		}
	}
	o.Stdoutln(rolls)
}

func (d *day04) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"",
			},
		},
		{
			ExpectedOutput: []string{
				"",
			},
		},
	}
}
