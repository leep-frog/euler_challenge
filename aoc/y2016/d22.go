package y2016

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/rgx"
)

func Day22() aoc.Day {
	return &day22{}
}

type day22 struct{}

type fileNode struct {
	x, y, used, available int
}

func (fn *fileNode) String() string {
	return fmt.Sprintf("(%d,%d, %d/%d)", fn.x, fn.y, fn.used, fn.available)
}

func (d *day22) Solve(lines []string, o command.Output) {
	r := rgx.New(`/dev/grid/node-x([0-9]+)-y([0-9]+)\s+([0-9]+)T\s+([0-9]+)T\s+([0-9]+)T\s+([0-9]+)%`)

	var nodes []*fileNode
	for _, line := range lines[2:] {
		m := r.MustMatch(line)
		nodes = append(nodes, &fileNode{
			parse.Atoi(m[0]), parse.Atoi(m[1]), parse.Atoi(m[3]), parse.Atoi(m[4]),
		})
	}

	// There is only empty space
	var grid [][]int
	sizeX, sizeY := nodes[len(nodes)-1].x+1, nodes[len(nodes)-1].y+1
	for i := 0; i < sizeX; i++ {
		var row []int
		for j := 0; j < sizeY; j++ {
			row = append(row, 0)
		}
		grid = append(grid, row)
	}

	var cnt int
	for i, a := range nodes {
		for j, b := range nodes {
			if i != j && a.used != 0 && a.used <= b.available {
				grid[b.x][b.y] = 2
				grid[a.x][a.y] = 1
				cnt++
			}
		}
	}

	// By viewing the grid printed out, you can see that the
	// solution just involves moving a single open node around
	// like a sliding puzzle.
	// Solution is (distance to corner) + (5 moves for each horizontal movement of goal node)
	//           =     (19 + 6 + 36)    + (5 * 35)
	//           = 236
	//
	parse.PrintAOCGrid(maths.SimpleTranspose(grid), map[int]rune{0: '#', 1: '.', 2: '_'})
	o.Stdoutln(cnt, 236)
}

/*
.....................................
.....................................
.....................................
.####################################
.....................................
.....................................
..................._.................
.....................................
.....................................
.....................................
.....................................
.....................................
.....................................
.....................................
.....................................
.....................................
.....................................
.....................................
.....................................
.....................................
.....................................
.....................................
.....................................
.....................................
.....................................
*/

func (d *day22) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"7 236",
			},
		},
		{
			ExpectedOutput: []string{
				"888 236",
			},
		},
	}
}
