package y2016

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/walker"
)

func Day13() aoc.Day {
	return &day13{}
}

type day13 struct{}

func (d *day13) Solve(lines []string, o command.Output) {
	var res []int
	for _, part2 := range []bool{false, true} {
		input := strings.Split(lines[0], ",")
		cc := &cubicleContext{parse.Atoi(input[0]), parse.Atoi(input[1]), parse.Atoi(input[2]), 0, part2}
		if part2 {
			cc.wantX = -1
		}
		_, dist := bfs.ContextPathSearch[string](cc, []*cubicle{{1, 1}})
		if part2 {
			res = append(res, cc.count)
		} else {
			res = append(res, dist)
		}
	}
	o.Stdoutln(res[0], res[1])
}

type cubicleContext struct {
	offset, wantX, wantY, count int
	part2                       bool
}

type cubicle struct {
	x, y int
}

func (c *cubicle) Done(ctx *cubicleContext, path bfs.Path[*cubicle]) bool {
	ctx.count++
	return c.x == ctx.wantX && c.y == ctx.wantY
}

func (c *cubicle) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c *cubicle) Code(ctx *cubicleContext, path bfs.Path[*cubicle]) string {
	return c.String()
}

func (c *cubicle) AdjacentStates(ctx *cubicleContext, path bfs.Path[*cubicle]) []*cubicle {
	if path.Len() > 50 && ctx.part2 {
		return nil
	}
	var neighbors []*cubicle
	for _, d := range walker.CardinalDirections(true) {
		newC := &cubicle{c.x + d.X, c.y + d.Y}
		if newC.x < 0 || newC.y < 0 {
			continue
		}
		v := newC.x*newC.x + 3*newC.x + 2*newC.x*newC.y + newC.y + newC.y*newC.y + ctx.offset
		if strings.Count(maths.ToBinary(v).String(), "1")%2 == 0 {
			neighbors = append(neighbors, newC)
		}
	}
	return neighbors
}

func (d *day13) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"11 151",
			},
		},
		{
			ExpectedOutput: []string{
				"92 124",
			},
		},
	}
}
