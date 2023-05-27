package y2017

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day21() aoc.Day {
	return &day21{}
}

type day21 struct{}

type conversions struct {
	input  [][]bool
	output [][]bool
}

func (c *conversions) convert(sg [][]bool) [][]bool {
	var matches bool
	for i := 0; i < 2 && !matches; i++ {
		if i == 1 {
			sg = maths.SimpleTranspose(sg)
		}
		for j := 0; j < 4 && !matches; j++ {
			if j != 0 {
				sg = maths.Rotate(sg)
			}
			if maths.MatrixEquals(c.input, sg) {
				matches = true
			}
		}
	}

	if !matches {
		return nil
	}
	return maths.MatrixCopy(c.output)
}

func (d *day21) Solve(lines []string, o command.Output) {
	parts := strings.Split(lines[0], ",")
	part1, part2 := parse.Atoi(parts[0]), parse.Atoi(parts[1])
	o.Stdoutln(d.solve(lines[1:], part1), d.solve(lines[1:], part2))
}

func (d *day21) solve(lines []string, iterations int) int {
	// .#.
	// ..#
	// ###
	grid := [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}

	var twoCnvrs, threeCnvrs []*conversions
	for _, line := range lines {
		parts := strings.Split(line, " => ")
		inputLines := functional.Map(strings.Split(parts[0], "/"), func(s string) []bool {
			return functional.Map(strings.Split(s, ""), func(c string) bool { return c == "#" })
		})
		outputLines := functional.Map(strings.Split(parts[1], "/"), func(s string) []bool {
			return functional.Map(strings.Split(s, ""), func(c string) bool { return c == "#" })
		})
		if len(inputLines) == 2 {
			twoCnvrs = append(twoCnvrs, &conversions{inputLines, outputLines})
		} else {
			threeCnvrs = append(threeCnvrs, &conversions{inputLines, outputLines})
		}
	}

	for i := 0; i < iterations; i++ {
		subGridSize := 3
		cnvrs := threeCnvrs
		if len(grid)%2 == 0 {
			subGridSize = 2
			cnvrs = twoCnvrs
		}

		// Create sub-grids
		var subGrids [][][][]bool
		for xStart := 0; xStart < len(grid); xStart += subGridSize {
			var subGridsRow [][][]bool
			for yStart := 0; yStart < len(grid); yStart += subGridSize {
				var subGrid [][]bool
				for x := xStart; x < xStart+subGridSize; x++ {
					var subRow []bool
					for y := yStart; y < yStart+subGridSize; y++ {
						subRow = append(subRow, grid[x][y])
					}
					subGrid = append(subGrid, subRow)
				}
				subGridsRow = append(subGridsRow, subGrid)
			}
			subGrids = append(subGrids, subGridsRow)
		}

		// Convert grids
		convertedSubGrids := functional.Map(subGrids, func(sgRows [][][]bool) [][][]bool {
			return functional.Map(sgRows, func(sg [][]bool) [][]bool {

				for _, c := range cnvrs {
					if res := c.convert(sg); len(res) != 0 {
						return res
					}
				}
				panic("No matching conversion")
			})
		})

		// Create new grid by stitching together sub grids
		var newGrid [][]bool
		for ri, csgr := range convertedSubGrids {
			for k := 0; k < subGridSize+1; k++ {
				newGrid = append(newGrid, []bool{})
			}
			for _, sg := range csgr {
				for i, r := range sg {
					newGrid[ri*(subGridSize+1)+i] = append(newGrid[ri*(subGridSize+1)+i], r...)
				}
			}
		}
		grid = newGrid
	}
	var cnt int
	for _, r := range grid {
		cnt += functional.Count(r, true)
	}

	return cnt
}

func (d *day21) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"4 12",
			},
		},
		{
			ExpectedOutput: []string{
				"164 2355110",
			},
		},
	}
}
