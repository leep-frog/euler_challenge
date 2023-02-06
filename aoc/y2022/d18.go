package y2022

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/point"
)

func Day18() aoc.Day {
	return &day18{}
}

type day18 struct{}

var (
	dropletMoves = []*point.Point3D{
		point.New3D(1, 0, 0),
		point.New3D(-1, 0, 0),
		point.New3D(0, 1, 0),
		point.New3D(0, -1, 0),
		point.New3D(0, 0, 1),
		point.New3D(0, 0, -1),
	}
)

func (d *day18) checkGrid(grid [][][]bool, point *point.Point3D) int {
	var count int
	for _, move := range dropletMoves {
		neighbor := point.Plus(move)
		if neighbor.X < 0 || neighbor.Y < 0 || neighbor.Z < 0 {
			count++
			continue
		}

		if neighbor.X >= len(grid) || neighbor.Y >= len(grid[0]) || neighbor.Z >= len(grid[0][0]) {
			count++
			continue
		}

		if !grid[neighbor.X][neighbor.Y][neighbor.Z] {
			count++
		}
	}
	return count
}

func (d *day18) Solve(lines []string, o command.Output) {
	droplets := functional.Map(parse.ToGrid(lines, ","), func(parts []int) *point.Point3D {
		return point.New3D(parts[0]+1, parts[1]+1, parts[2]+1)
	})

	var maxX, maxY, maxZ int
	for _, d := range droplets {
		maxX = maths.Max(maxX, d.X)
		maxY = maths.Max(maxY, d.Y)
		maxZ = maths.Max(maxZ, d.Z)
	}

	var grid [][][]bool
	for x := 0; x <= maxX+2; x++ {
		var twoD [][]bool
		for y := 0; y <= maxY+2; y++ {
			twoD = append(twoD, make([]bool, maxZ+3))
		}
		grid = append(grid, twoD)
	}

	for _, droplet := range droplets {
		grid[droplet.X][droplet.Y][droplet.Z] = true
	}

	var count int
	for _, droplet := range droplets {
		count += d.checkGrid(grid, droplet)
	}

	ctx := &dropletCtx{grid, 0}
	start := &dropletNode{point.Origin3D()}
	bfs.ContextSearch[string](ctx, []*dropletNode{start})
	o.Stdoutln(count, ctx.count)
}

type dropletNode struct {
	p *point.Point3D
}

type dropletCtx struct {
	grid  [][][]bool
	count int
}

func (d *dropletNode) Code(ctx *dropletCtx) string {
	return d.p.String()
}

func (d *dropletNode) Done(ctx *dropletCtx) bool {
	return false
}

func (d *dropletNode) AdjacentStates(ctx *dropletCtx) []*dropletNode {
	if ctx.grid[d.p.X][d.p.Y][d.p.Z] {
		return nil
	}

	var r []*dropletNode
	for _, move := range dropletMoves {
		neighbor := d.p.Plus(move)
		if neighbor.X < 0 || neighbor.Y < 0 || neighbor.Z < 0 {
			continue
		}

		if neighbor.X >= len(ctx.grid) || neighbor.Y >= len(ctx.grid[0]) || neighbor.Z >= len(ctx.grid[0][0]) {
			continue
		}

		if ctx.grid[neighbor.X][neighbor.Y][neighbor.Z] {
			ctx.count++
		} else {
			r = append(r, &dropletNode{neighbor})
		}
	}

	return r
}

func (d *day18) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"64 58",
			},
		},
		{
			ExpectedOutput: []string{
				"4242 2428",
			},
		},
	}
}
