package eulerchallenge

import (
	"math/big"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/maths"
)

func P151() *problem {
	return intInputNode(151, func(o command.Output, n int) {

		m := map[int]int{}
		for j := 2; j <= n; j++ {
			m[j] = 1
		}

		ctx := &context151{n, map[int]*big.Rat{}}
		init := []*node151{{m, n - 1, big.NewRat(1, 1)}}
		bfs.ContextualDFS(init, ctx, bfs.AllowDFSCycles(), bfs.AllowDFSDuplicates())
		sm := big.NewRat(0, 1)
		for i := 0; i < maths.Pow(2, n-1)-1; i++ {
			if num := ctx.dayFreq[i]; num != nil {
				sm = big.NewRat(1, 1).Add(sm, num)
			}
		}
		f, _ := sm.Float64()
		o.Stdoutf("%.6f\n", f)
	}, []*execution{
		{
			args:     []string{"5"},
			want:     "0.464399",
			estimate: 0.5,
		},
		{
			args: []string{"3"},
			want: "0.500000",
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
	freq *big.Rat
}

type context151 struct {
	n int
	// map from day to how often that day had one piece of paper.
	dayFreq map[int]*big.Rat
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
			ctx.dayFreq[dp.Len()] = ratAdd(ctx.dayFreq[dp.Len()], n.freq)
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

		r = append(r, &node151{newSizes, n.numSheets - 1 + (ctx.n - k), ratMul(newRat(v, n.numSheets), n.freq)})
	}
	return r
}
