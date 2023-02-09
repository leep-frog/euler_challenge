package y2022

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day08() aoc.Day {
	return &day08{}
}

type day08 struct{}

func (d *day08) visibleCount(grid [][]int, inverted bool, visible [][]string) { //m map[int]map[int]bool) {
	for x, row := range grid {
		leftHeight := -1
		rightHeight := -1
		for y, fromLeft := range row {
			fromRight := row[len(row)-1-y]

			if fromLeft > leftHeight {
				if !inverted {
					visible[x][y] = "L"
					// maths.Insert(m, x, y, true)
				} else {
					visible[y][x] = "U"
					// maths.Insert(m, y, x, true)
				}
				leftHeight = fromLeft
			}

			if fromRight > rightHeight {
				if !inverted {
					visible[x][len(row)-1-y] = "R"
					// maths.Insert(m, len(row)-1-x, y, true)
				} else {
					visible[len(row)-1-y][x] = "D"
					// maths.Insert(m, y, len(row)-1-x, true)
				}
				rightHeight = fromRight
			}
		}
	}
}

func (d *day08) scenicScores(grid [][]int, x, y int) int {
	h := grid[x][y]

	// left
	var left int
	for i := x - 1; ; i-- {
		if i < 0 {
			break
		}
		left++
		if grid[i][y] >= h {
			break
		}
	}

	// right
	var right int
	for i := x + 1; ; i++ {
		if i >= len(grid) {
			break
		}
		right++
		if grid[i][y] >= h {
			break
		}
	}

	// down
	var down int
	for j := y - 1; ; j-- {
		if j < 0 {
			break
		}
		down++
		if grid[x][j] >= h {
			break
		}
	}

	// up
	var up int
	for j := y + 1; ; j++ {
		if j >= len(grid[0]) {
			break
		}
		up++
		if grid[x][j] >= h {
			break
		}
	}

	return left * right * down * up
}

func (d *day08) Solve(lines []string, o command.Output) {
	grid := parse.ToGrid(lines, "")

	var visible [][]string
	for _, row := range grid {
		var vr []string
		for _ = range row {
			vr = append(vr, "_")
		}
		visible = append(visible, vr)
	}

	d.visibleCount(grid, false, visible)
	d.visibleCount(maths.SimpleTranspose(grid), true, visible)

	var sum int
	for _, row := range visible {
		for _, c := range row {
			if c != "_" {
				sum++
			}
		}
	}

	best := maths.Largest[int, int]()
	for x, row := range grid {
		for y := range row {
			best.Check(d.scenicScores(grid, x, y))
		}
	}
	o.Stdoutln(sum, best.Best())
}

func (d *day08) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"21 8",
			},
		},
		{
			ExpectedOutput: []string{
				"1787 440640",
			},
		},
	}
}
