package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"

	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/parse"
)

func P82() *problem {
	return fileInputNode(82, func(lines []string, o command.Output) {
		grid := parse.ToGrid(lines)
		var initStates []*p82
		for i := 0; i < len(grid); i++ {
			initStates = append(initStates, &p82{i, 0})
		}
		_, dist := bfs.ContextualShortestOffsetPath[[][]int](initStates, grid)
		o.Stdoutln(dist)
	})
}

type p82 struct {
	i, j int
}

func (p *p82) Code(grid [][]int) string {
	return fmt.Sprintf("%d_%d", p.i, p.j)
}

func (p *p82) Distance(grid [][]int) int {
	return grid[p.i][p.j]
}

func (p *p82) Done(grid [][]int) bool {
	return p.j == len(grid[p.i])-1
}

func (p *p82) AdjacentStates(grid [][]int) []*p82 {
	r := []*p82{}
	if p.i < len(grid)-1 {
		r = append(r, &p82{p.i + 1, p.j})
	}
	if p.j < len(grid[p.i])-1 {
		r = append(r, &p82{p.i, p.j + 1})
	}
	if p.i > 0 {
		r = append(r, &p82{p.i - 1, p.j})
	}
	return r
}
