package eulerchallenge

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
)

type grid166 struct {
	values    [][]int
	populated [][]bool
	sum       int
	d1Sum     []int
	d1Rem     []int
	d2Sum     []int
	d2Rem     []int
	diag1Sum  int
	diag1Rem  int
	diag2Sum  int
	diag2Rem  int
}

func newGrid166(sum int) *grid166 {
	var vs [][]int
	var populated [][]bool
	var d1Rem, d2Rem []int
	for i := 0; i < 4; i++ {
		d1Rem = append(d1Rem, 4)
		d2Rem = append(d2Rem, 4)
		var vr []int
		var pr []bool
		for j := 0; j < 4; j++ {
			vr = append(vr, 0)
			pr = append(pr, false)
		}
		vs = append(vs, vr)
		populated = append(populated, pr)
	}

	return &grid166{vs, populated, sum, make([]int, 4, 4), d1Rem, make([]int, 4, 4), d2Rem, 0, 4, 0, 4}
}

func (g *grid166) String() string {
	var sl []string
	for i := 0; i < 4; i++ {
		var sr []string
		for j := 0; j < 4; j++ {
			if g.populated[j][i] {
				sr = append(sr, fmt.Sprintf("%d", g.values[j][i]))
			} else {
				sr = append(sr, "X")
			}
		}
		sl = append(sl, strings.Join(sr, " "))
	}
	return strings.Join(sl, "\n") + fmt.Sprintf(", Sum = %d", g.sum)
}

func (g *grid166) set(p *point.Point[int], v int) {
	g.populated[p.X][p.Y] = true
	g.values[p.X][p.Y] = v

	g.d1Rem[p.X]--
	g.d1Sum[p.X] += v

	g.d2Rem[p.Y]--
	g.d2Sum[p.Y] += v

	if p.X == p.Y {
		g.diag1Rem--
		g.diag1Sum += v
	}
	if p.X+p.Y == 3 {
		g.diag2Rem--
		g.diag2Sum += v
	}
}

func (g *grid166) unset(p *point.Point[int]) {
	g.populated[p.X][p.Y] = false
	v := g.values[p.X][p.Y]
	g.values[p.X][p.Y] = 0

	g.d1Rem[p.X]++
	g.d1Sum[p.X] -= v

	g.d2Rem[p.Y]++
	g.d2Sum[p.Y] -= v

	if p.X == p.Y {
		g.diag1Rem++
		g.diag1Sum -= v
	}
	if p.X+p.Y == 3 {
		g.diag2Rem++
		g.diag2Sum -= v
	}
}

func (g *grid166) d1State(x int) (int, int) {
	return g.d1Rem[x], g.d1Sum[x]
}

func (g *grid166) d2State(y int) (int, int) {
	return g.d2Rem[y], g.d2Sum[y]
}

func (g *grid166) diag1State() (int, int) {
	return g.diag1Rem, g.diag1Sum
}

func (g *grid166) diag2State() (int, int) {
	return g.diag2Rem, g.diag2Sum
}

// TODO: make point.Grid (that takes point as index)
func (g *grid166) minMaxValues(p *point.Point[int]) (int, int) {
	// Row value
	need := g.sum - g.d1Sum[p.X]
	min, max := maths.Max(0, need-((g.d1Rem[p.X]-1)*9)), maths.Min(9, need)

	// Col value
	need = g.sum - g.d2Sum[p.Y]
	min, max = maths.Max(min, need-((g.d2Rem[p.Y]-1)*9)), maths.Min(max, need)

	// Diagonals
	if p.X == p.Y {
		need = g.sum - g.diag1Sum
		min, max = maths.Max(min, need-((g.diag1Rem-1)*9)), maths.Min(max, need)
	}

	if p.X+p.Y == 3 {
		need = g.sum - g.diag2Sum
		min, max = maths.Max(min, need-((g.diag2Rem-1)*9)), maths.Min(max, need)
	}

	return min, max
}

var (
	iterOrder = []*point.Point[int]{
		point.New(0, 0),
		point.New(1, 1),
		point.New(2, 2),
		point.New(3, 2),
		point.New(3, 1),
		point.New(0, 3),
		point.New(1, 2),
		point.New(1, 0),
		// Deducable:
		point.New(3, 3),
		point.New(3, 0),
		point.New(2, 0),
		point.New(2, 1),
		point.New(0, 1),
		point.New(0, 2),
		point.New(1, 3),
		point.New(2, 3),
	}
)

func rec166(idx int, grid *grid166) int {
	if idx == 16 {
		return 1
	}

	pt := iterOrder[idx]

	minV, maxV := grid.minMaxValues(pt)

	// Still iterating over cells with multiple options
	if idx <= 7 {
		var cnt int
		for v := minV; v <= maxV; v++ {
			grid.set(pt, v)
			cnt += rec166(idx+1, grid)
			grid.unset(pt)
		}
		return cnt
	}

	// Otherwise, we have none or exactly one option for the cell.
	if minV != maxV {
		return 0
	}
	grid.set(pt, minV)
	c := rec166(idx+1, grid)
	grid.unset(pt)
	return c

}

func P166() *problem {
	return noInputNode(166, func(o command.Output) {
		c := 0
		for sum := 0; sum <= 18; sum++ {
			got := rec166(0, newGrid166(sum))
			if sum == 18 {
				c += got
			} else {
				c += 2 * got
			}
		}
		o.Stdoutln(c)
	}, &execution{
		want:     "7130034",
		estimate: 10,
	})
}
