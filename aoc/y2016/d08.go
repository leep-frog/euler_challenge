package y2016

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day08() aoc.Day {
	return &day08{}
}

type day08 struct{}

func (d *day08) rotate(grid [][]bool, row, by int) [][]bool {
	grid[row] = append(grid[row][len(grid[row])-by:], grid[row][:len(grid[row])-by]...)
	return grid
}

func (d *day08) rotateCol(grid [][]bool, row, by int) [][]bool {
	grid = maths.SimpleTranspose(grid)
	return maths.SimpleTranspose(d.rotate(grid, row, by))
}

func (d *day08) Solve(lines []string, o command.Output) {
	var grid [][]bool
	rows, cols := 6, 50
	// rows, cols := 3, 7
	for i := 0; i < rows; i++ {
		grid = append(grid, make([]bool, cols, cols))
	}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "rect":
			dims := strings.Split(parts[1], "x")
			J, I := parse.Atoi(dims[0]), parse.Atoi(dims[1])
			for i := 0; i < I; i++ {
				for j := 0; j < J; j++ {
					grid[i][j] = true
				}
			}
		case "rotate":
			idx := parse.Atoi(strings.Split(parts[2], "=")[1])
			by := parse.Atoi(parts[4])
			if parts[1] == "row" {
				grid = d.rotate(grid, idx, by)
			} else if parts[1] == "column" {
				grid = d.rotateCol(grid, idx, by)
			} else {
				panic("UGH")
			}
		}
	}

	o.Stdoutln(functional.Count2D(grid, true))

	// Uncomment the below line to solve part 2
	// parse.PrintAOCGrid(grid, map[bool]rune{true: '#', false: ' '})
}
func (d *day08) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"6",
			},
		},
		{
			ExpectedOutput: []string{
				"110",
				// ZJHRKCPLYJ
			},
		},
	}
}
