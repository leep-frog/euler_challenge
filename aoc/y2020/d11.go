package y2020

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day11() aoc.Day {
	return &day11{}
}

type day11 struct{}

type seat int

const (
	floor seat = iota
	emptySeat
	filledSeat
)

func (d *day11) occupiedCount(grid [][]seat, i, j int, part2 bool) int {
	moves := [][]int{
		{0, 1},
		{0, -1},
		{1, 1},
		{1, 0},
		{1, -1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
	}

	var count int
	for _, m := range moves {
		for x, y := i+m[0], j+m[1]; x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]); x, y = x+m[0], y+m[1] {
			if grid[x][y] == filledSeat {
				count++
				break
			}
			if !part2 || grid[x][y] == emptySeat {
				break
			}
		}
	}
	return count
}

func (d *day11) solve(lines []string, threshold int, part2 bool) int {
	grid := parse.MapToGrid(lines, map[rune]seat{
		'.': floor,
		'L': emptySeat,
	})

	var seatCount int
	for change := true; change; {
		change = false
		var newGrid [][]seat
		for i, row := range grid {
			var newRow []seat
			for j := range row {
				switch grid[i][j] {
				case floor:
					newRow = append(newRow, floor)
				case emptySeat:
					if d.occupiedCount(grid, i, j, part2) == 0 {
						newRow = append(newRow, filledSeat)
						seatCount++
						change = true
					} else {
						newRow = append(newRow, emptySeat)
					}
				case filledSeat:
					if d.occupiedCount(grid, i, j, part2) >= threshold {
						newRow = append(newRow, emptySeat)
						seatCount--
						change = true
					} else {
						newRow = append(newRow, filledSeat)
					}
				}
			}
			newGrid = append(newGrid, newRow)
		}
		grid = newGrid
	}
	return seatCount
}

func (d *day11) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines, 4, false), d.solve(lines, 5, true))
}

func (d *day11) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"37 26",
			},
		},
		{
			ExpectedOutput: []string{
				"2247 2011",
			},
		},
	}
}
