package eulerchallenge

import (
	"fmt"
	"math"
	"sort"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/maths"
)

func P119() *problem {
	return intInputNode(119, func(o command.Output, n int) {
		ctx := &context119{map[string]bool{}, n, nil}
		bfs.ShortestWeightedPath([]*node119{{2, 2, maths.NewInt(4)}}, ctx)

		var values []string
		for k := range ctx.values {
			values = append(values, k)
		}
		sort.SliceStable(values, func(i, j int) bool {
			if len(values[i]) != len(values[j]) {
				return len(values[i]) < len(values[j])
			}
			return values[i] < values[j]
		})
		o.Stdoutln(values[n-1])
	})
}

type node119 struct {
	int_ int
	pow int
	p *maths.Int
}

type context119 struct {
	values map[string]bool
	n int
	max *maths.Int
}

func (n *node119) Code(ctx *bfs.Context[*context119, *node119]) string {
	return fmt.Sprintf("%v^%d", n.int_, n.pow)
}

func (n *node119) Distance(ctx *bfs.Context[*context119, *node119]) int {
	return int(float64(n.pow)*math.Log(float64(n.int_)))
}

func (n *node119) Done(ctx *bfs.Context[*context119, *node119]) bool {
	if maths.SumSys(n.p.Digits()...) == n.int_ {
		ctx.GlobalContext.values[n.p.String()] = true
		if ctx.GlobalContext.max == nil || n.p.GT(ctx.GlobalContext.max) {
			ctx.GlobalContext.max = n.p
		}
	}
	return false
}

func (n *node119) AdjacentStates(ctx *bfs.Context[*context119, *node119]) []*node119 {	
	if len(ctx.GlobalContext.values) >= ctx.GlobalContext.n && n.p.GT(ctx.GlobalContext.max) {
		return nil
	}

	// Squared numbers over 100 can't work because 
	// max digit sum of n < n; when n > 100
	// 10*log_10(n) < n; when n > 100
	if n.pow == 2 && n.int_ > 100 {
		return nil
	}
	return []*node119{
		{n.int_, n.pow+1, maths.BigPow(n.int_, n.pow+1)},
		{n.int_+1, n.pow, maths.BigPow(n.int_+1, n.pow)},
	}
}