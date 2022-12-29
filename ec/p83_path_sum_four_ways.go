package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"

	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/parse"
)

func P83() *problem {
	return fileInputNode(83, func(lines []string, o command.Output) {
		_, dist := bfs.ContextDistanceSearch[[][]int, string, bfs.Int](parse.ToGrid(lines, ","), []*p83{{}})
		o.Stdoutln(dist)
	}, []*execution{
		{
			args: []string{"p83.txt"},
			want: "425185",
		},
	})
}

type p83 struct {
	i, j int
}

func (p *p83) Code(grid [][]int) string {
	return fmt.Sprintf("%d_%d", p.i, p.j)
}

func (p *p83) Distance(grid [][]int) bfs.Int {
	return bfs.Int(grid[p.i][p.j])
}

func (p *p83) Done(grid [][]int) bool {
	return p.i == len(grid)-1 && p.j == len(grid[p.i])-1
}

func (p *p83) AdjacentStates(grid [][]int) []*p83 {
	r := []*p83{}
	if p.i < len(grid)-1 {
		r = append(r, &p83{p.i + 1, p.j})
	}
	if p.j < len(grid[p.i])-1 {
		r = append(r, &p83{p.i, p.j + 1})
	}
	if p.i > 0 {
		r = append(r, &p83{p.i - 1, p.j})
	}
	if p.j > 0 {
		r = append(r, &p83{p.i, p.j - 1})
	}
	return r
}
