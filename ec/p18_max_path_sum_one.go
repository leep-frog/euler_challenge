package eulerchallenge

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

var (
	maxValue = 100
)

func P18() *command.Node {
	return command.SerialNodes(
		command.Description(""),
		command.StringNode("FILE", ""),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			lines := parse.ReadFileLines(d.String("FILE"))

			var tower [][]int
			for _, line := range lines {
				var row []int
				for _, c := range strings.Split(line, " ") {
					row = append(row, parse.Atoi(c))
				}
				tower = append(tower, row)
			}

			path, dist := bfs.ShortestOffsetPath(&place{0,0}, maxValue-tower[0][0], tower)
			for _, p := range path {
				o.Stdoutln(p)
			}
			o.Stdoutln((maxValue * len(tower)) - dist)
		}),
	)
}

func check(tower [][]int, row, col, sum int) int {
	if row == len(tower)-1 {
		return sum
	}

	left := check(tower, row+1, col, sum+tower[row+1][col])
	right := check(tower, row+1, col+1, sum+tower[row+1][col+1])
	return maths.Max(left, right)
}

type place struct {
	row int
	col int
}

func (p *place) String() string {
	return p.Code()
}

func (p *place) Code() string {
	return fmt.Sprintf("%d_%d", p.row, p.col)
}

func (p *place) Done(tower [][]int, dist int) bool {
	return p.row == len(tower)-1
}

func (p *place) AdjacentStates(tower [][]int) []*bfs.AdjacentState[[][]int, *place] {
	return []*bfs.AdjacentState[[][]int, *place]{
		{
			State: &place{
				col: p.col,
				row: p.row + 1,
			},
			Offset: maxValue - tower[p.row+1][p.col],
		},
		{
			State: &place{
				col: p.col + 1,
				row: p.row + 1,
			},
			Offset: maxValue - tower[p.row+1][p.col+1],
		},
	}
}
