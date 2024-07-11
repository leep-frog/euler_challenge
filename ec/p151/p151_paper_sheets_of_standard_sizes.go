package p151

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/maths"
)

func P151() *ecmodels.Problem {
	return ecmodels.IntInputNode(151, func(o command.Output, n int) {

		m := map[int]int{}
		for j := 2; j <= n; j++ {
			m[j] = 1
		}

		ctx := &context151{n, map[int]*fraction.Rational{}}
		init := []*node151{{m, n - 1, fraction.NewRational(1, 1)}}
		bfs.ContextualDFS(init, ctx, bfs.AllowDFSCycles(), bfs.AllowDFSDuplicates())
		sm := fraction.NewRational(0, 1)
		for i := 0; i < maths.Pow(2, n-1)-1; i++ {
			if num := ctx.dayFreq[i]; num != nil {
				sm = sm.Plus(num)
			}
		}
		o.Stdoutf("%.6f\n", sm.Float64())
	}, []*ecmodels.Execution{
		{
			Args:     []string{"5"},
			Want:     "0.464399",
			Estimate: 0.5,
		},
		{
			Args: []string{"3"},
			Want: "0.500000",
		},
	})
}

type node151 struct {
	// map from size to count of that size
	sizes map[int]int
	// number of sheets in the bag
	numSheets int
	// probability we get to this configuration
	// note we may get to the same configuration two different ways
	// and that would be represented by separate node151 instances.
	freq *fraction.Rational
}

type context151 struct {
	n int
	// map from day to how often that day had one piece of paper.
	dayFreq map[int]*fraction.Rational
}

func (n *node151) Code(*context151, bfs.DFSPath[*node151]) string { return "" }
func (n *node151) Done(ctx *context151, dp bfs.DFSPath[*node151]) bool {
	if len(n.sizes) == 0 {
		return false
	}
	if n.numSheets == 1 {
		if ctx.dayFreq[dp.Len()] == nil {
			ctx.dayFreq[dp.Len()] = n.freq
		} else {
			ctx.dayFreq[dp.Len()] = ctx.dayFreq[dp.Len()].Plus(n.freq)
		}
	}
	return false
}

func (n *node151) AdjacentStates(ctx *context151, dp bfs.DFSPath[*node151]) []*node151 {
	var r []*node151
	for k, v := range n.sizes {
		if v == 0 {
			continue
		}
		newSizes := maths.CopyMap(n.sizes)
		if v == 1 {
			delete(newSizes, k)
		} else {
			newSizes[k]--
		}

		for paper := k + 1; paper <= ctx.n; paper++ {
			newSizes[paper]++
		}

		r = append(r, &node151{newSizes, n.numSheets - 1 + (ctx.n - k), fraction.NewRational(v, n.numSheets).Times(n.freq)})
	}
	return r
}
