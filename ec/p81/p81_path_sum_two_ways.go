package p81

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/parse"
)

func P81() *ecmodels.Problem {
	return ecmodels.FileInputNode(81, func(lines []string, o command.Output) {
		_, dist := bfs.ContextDistanceSearch[string, bfs.Int](parse.ToGrid(lines, ","), []*p81{{}})
		o.Stdoutln(dist)
	}, []*ecmodels.Execution{
		{
			Args: []string{"p81_example.txt"},
			Want: "2427",
		},
		{
			Args: []string{"p81.txt"},
			Want: "427337",
		},
	})
}

type p81 struct {
	i, j int
}

func (p *p81) Code([][]int) string {
	return fmt.Sprintf("%d_%d", p.i, p.j)
}

func (p *p81) Distance(grid [][]int) bfs.Int {
	return bfs.Int(grid[p.i][p.j])
}

func (p *p81) Done(grid [][]int) bool {
	return p.i == len(grid)-1 && p.j == len(grid[p.i])-1
}

func (p *p81) AdjacentStates(grid [][]int) []*p81 {
	r := []*p81{}
	if p.i < len(grid)-1 {
		r = append(r, &p81{p.i + 1, p.j})
	}
	if p.j < len(grid[p.i])-1 {
		r = append(r, &p81{p.i, p.j + 1})
	}
	return r
}
