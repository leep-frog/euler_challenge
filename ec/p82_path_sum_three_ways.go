package eulerchallenge

import (
	"github.com/leep-frog/command"
	"fmt"

	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/bfs"
)

func P82() *problem {
	return fileInputNode(82, func(lines []string, o command.Output) {
		grid := parse.ToGrid(lines)
		var initStates []*p82
		for i := 0; i < len(grid); i++ {
			initStates = append(initStates, &p82{i, 0})
		}
			_, dist := bfs.ShortestOffsetPath[[][]int, *p82](initStates, grid)
			o.Stdoutln(dist)
	})
}

type p82 struct {
	i, j int
}

func (p *p82) Code(ctx *bfs.Context[[][]int, *p82]) string {
	return fmt.Sprintf("%d_%d", p.i, p.j)
}

func (p *p82) Offset(ctx *bfs.Context[[][]int, *p82]) int {
	grid := ctx.GlobalContext
	return grid[p.i][p.j]
}

func (p *p82) Done(ctx *bfs.Context[[][]int, *p82]) bool {
	return p.j == len(ctx.GlobalContext[p.i]) - 1
}

func (p *p82) AdjacentStates(ctx *bfs.Context[[][]int, *p82]) []*p82 {
	grid := ctx.GlobalContext
	r := []*p82{}
	if p.i < len(grid) - 1 {
		r = append(r, &p82{p.i+1, p.j})
	}
	if p.j < len(grid[p.i]) - 1 {
		r = append(r, &p82{p.i, p.j+1})
	}
	if p.i > 0 {
		r = append(r, &p82{p.i-1, p.j})
	}
	return r
}