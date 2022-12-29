package y2022

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/bfs"
)

func Day12() aoc.Day {
	return &day12{}
}

type day12 struct{}

type context12 struct {
	grid [][]int
	endX int
	endY int
}

// TODO: bfs
type height struct {
	x, y int
}

func (h *height) String() string {
	return fmt.Sprintf("(%d, %d)", h.x, h.y)
}

func (h *height) Code(ctx *context12) string {
	return fmt.Sprintln(h)
}

func (h *height) Done(ctx *context12) bool {
	return h.x == ctx.endX && h.y == ctx.endY
}

func (h *height) AdjacentStates(ctx *context12) []*height {
	v := ctx.grid[h.x][h.y]
	var r []*height
	if h.x > 0 && ctx.grid[h.x-1][h.y] <= v+1 {
		r = append(r, &height{h.x - 1, h.y})
	}
	if h.y > 0 && ctx.grid[h.x][h.y-1] <= v+1 {
		r = append(r, &height{h.x, h.y - 1})
	}
	if h.x < len(ctx.grid)-1 && ctx.grid[h.x+1][h.y] <= v+1 {
		r = append(r, &height{h.x + 1, h.y})
	}
	if h.y < len(ctx.grid[0])-1 && ctx.grid[h.x][h.y+1] <= v+1 {
		r = append(r, &height{h.x, h.y + 1})
	}
	return r
}

func (d *day12) Solve(lines []string, o command.Output) {
	ctx := &context12{}
	var start *height
	var bases []*height
	for x, line := range lines {
		var row []int
		for y, c := range line {
			if c == 'S' {
				row = append(row, 0)
				start = &height{x, y}
				bases = append(bases, &height{x, y})
			} else if c == 'E' {
				row = append(row, 25)
				ctx.endX, ctx.endY = x, y
			} else {
				row = append(row, int(c-'a'))
				if c == 'a' {
					bases = append(bases, &height{x, y})
				}
			}
		}
		ctx.grid = append(ctx.grid, row)
	}

	_, dist1 := bfs.ContextSearch[*context12, string](ctx, []*height{start})
	_, dist2 := bfs.ContextSearch[*context12, string](ctx, bases)
	o.Stdoutln(dist1, dist2)
}

func (d *day12) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"31 29",
			},
		},
		{
			ExpectedOutput: []string{
				"412 402",
			},
		},
	}
}
