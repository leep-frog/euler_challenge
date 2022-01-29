package eulerchallenge

import (
	"fmt"
  "github.com/leep-frog/command"

	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/parse"
)

func P83() *problem {
	return fileInputNode(83, func(lines []string, o command.Output) {
		_, dist := bfs.ShortestOffsetPath[[][]int, *p83]([]*p83{{}}, parse.ToGrid(lines))
		o.Stdoutln(dist)
	})
}

type p83 struct {
	i, j int
}

func (p *p83) Code(ctx *bfs.Context[[][]int, *p83]) string {
	return fmt.Sprintf("%d_%d", p.i, p.j)
}

func (p *p83) Offset(ctx *bfs.Context[[][]int, *p83]) int {
	grid := ctx.GlobalContext
	return grid[p.i][p.j]
}

func (p *p83) Done(ctx *bfs.Context[[][]int, *p83]) bool {
	return p.i == len(ctx.GlobalContext) - 1 && p.j == len(ctx.GlobalContext[p.i]) - 1
}

func (p *p83) AdjacentStates(ctx *bfs.Context[[][]int, *p83]) []*p83 {
	grid := ctx.GlobalContext
	r := []*p83{}
	if p.i < len(grid) - 1 {
		r = append(r, &p83{p.i+1, p.j})
	}
	if p.j < len(grid[p.i]) - 1 {
		r = append(r, &p83{p.i, p.j+1})
	}
	if p.i > 0 {
		r = append(r, &p83{p.i-1, p.j})
	}
	if p.j > 0 {
		r = append(r, &p83{p.i, p.j-1})
	}
	return r
}