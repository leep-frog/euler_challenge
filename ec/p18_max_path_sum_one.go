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

func P18() *problem {
	return fileInputNode(18, func(lines []string, o command.Output) {
		var tower [][]int
		for _, line := range lines {
			var row []int
			for _, c := range strings.Split(line, " ") {
				row = append(row, parse.Atoi(c))
			}
			tower = append(tower, row)
		}

		_, dist := bfs.ContextDistanceSearch[[][]int, string, bfs.Int](tower, []*place{{0, 0}})
		o.Stdoutln((maxValue * len(tower)) - int(dist))
	}, []*execution{
		{
			args: []string{"p18_example.txt"},
			want: "23",
		},
		{
			args: []string{"p18.txt"},
			want: "1074",
		},
		{
			args: []string{"p67.txt"},
			want: "7273",
		},
	})
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
	return fmt.Sprintf("%d_%d", p.row, p.col)
}

func (p *place) Code([][]int) string {
	return p.String()
}

func (p *place) Done(tower [][]int) bool {
	return p.row == len(tower)-1
}

func (p *place) Distance(tower [][]int) bfs.Int {
	return bfs.Int(maxValue - tower[p.row][p.col])
}

func (p *place) AdjacentStates([][]int) []*place {
	return []*place{
		{
			col: p.col,
			row: p.row + 1,
		},
		{
			col: p.col + 1,
			row: p.row + 1,
		},
	}
}
