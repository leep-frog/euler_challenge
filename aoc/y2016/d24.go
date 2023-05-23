package y2016

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/walker"
	"golang.org/x/exp/maps"
)

func Day24() aoc.Day {
	return &day24{}
}

type day24 struct{}

type duct struct {
	x, y int
}

type ductContext struct {
	start     int
	grid      [][]int
	distances map[int]map[int]int
}

type ductPath struct {
	remaining map[int]bool
	path      []int
	loc, dist int
	part2     bool
}

func (dp *ductPath) Code(map[int]map[int]int) string {
	r := strings.Join(functional.Map(dp.path, func(k int) string { return fmt.Sprintf("%d", k) }), ",")
	return r
}

func (dp *ductPath) String(map[int]map[int]int) string {
	return dp.Code(nil)
}

func (dp *ductPath) Done(map[int]map[int]int) bool {
	return len(dp.remaining) == 0
}

func (dp *ductPath) Distance(distances map[int]map[int]int) bfs.Int {
	if dp.part2 && dp.Done(distances) {
		return bfs.Int(dp.dist + distances[dp.loc][0])
	}
	return bfs.Int(dp.dist)
}

func (dp *ductPath) AdjacentStates(distances map[int]map[int]int) []*ductPath {
	var r []*ductPath
	for k := range dp.remaining {
		newDP := &ductPath{
			maps.Clone(dp.remaining),
			append(bread.Copy(dp.path), k),
			k,
			dp.dist + distances[dp.loc][k],
			dp.part2,
		}

		delete(newDP.remaining, k)
		r = append(r, newDP)
	}
	return r
}

type ductPathContext struct {
	distances map[int]map[int]int
	remaining map[int]bool
}

func (d *duct) Code(*ductContext, bfs.Path[*duct]) string {
	return fmt.Sprintf("(%d,%d)", d.x, d.y)
}

func (d *duct) Done(ctx *ductContext, path bfs.Path[*duct]) bool {
	if v := ctx.grid[d.x][d.y]; v >= 0 {
		maths.Insert(ctx.distances, ctx.start, v, path.Len()-1)
		maths.Insert(ctx.distances, v, ctx.start, path.Len()-1)
	}
	return false
}

func (d *duct) AdjacentStates(ctx *ductContext, path bfs.Path[*duct]) []*duct {
	var r []*duct
	for _, dir := range walker.CardinalDirections(true) {
		newD := &duct{d.x + dir.X, d.y + dir.Y}
		if newD.x < 0 || newD.x >= len(ctx.grid) || newD.y < 0 || newD.y >= len(ctx.grid[0]) || ctx.grid[newD.x][newD.y] == -2 {
			continue
		}
		r = append(r, newD)
	}
	return r
}

func (d *day24) Solve(lines []string, o command.Output) {
	var grid [][]int
	var max int
	positions := map[int]*duct{}
	rem := map[int]bool{0: true}
	for x, line := range lines {
		var row []int
		for y, c := range line {
			if c == '.' {
				row = append(row, -1)
			} else if c == '#' {
				row = append(row, -2)
			} else {
				v := parse.Atoi(string(c))
				rem[v] = true
				positions[v] = &duct{x, y}
				row = append(row, v)
				if v > max {
					max = v
				}
			}
		}
		grid = append(grid, row)
	}

	distances := map[int]map[int]int{}
	for i := 0; i < max; i++ {
		bfs.ContextPathSearch[string](&ductContext{i, grid, distances}, []*duct{positions[i]})
	}
	o.Stdoutln(d.solve(distances, false), d.solve(distances, true))
}

func (d *day24) solve(distances map[int]map[int]int, part2 bool) int {
	rem := map[int]bool{}
	for i := 1; i < len(distances); i++ {
		rem[i] = true
	}

	_, dist := bfs.ContextDistanceSearch[string, bfs.Int](distances, []*ductPath{{rem, []int{0}, 0, 0, part2}}, bfs.CumulativeDistanceFunction())
	return int(dist)
}

func (d *day24) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"14 20",
			},
		},
		{
			ExpectedOutput: []string{
				"470 720",
			},
		},
	}
}
