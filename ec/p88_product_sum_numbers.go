package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/generator"
)

func P88() *problem {
	return intInputNode(88, func(o command.Output, n int) {
		g := generator.Primes()

		kMap := map[int]int{}
		// Max solution for a given k is 2k:
		// 2k = 1 * 1 * ... * 1 * 1 * 2 * k
		// sum = (k - 2)*1 + 2 + k = k - 2 + k = 2k
		for i := 2; i <= 2*n; i++ {
			ctx := &ctx88{nil, i, kMap, g}
			bfs.PoppableContextualDFS([]*n88{{i, 0, 0}}, ctx)
		}

		var sum int
		unique := map[int]bool{}
		for k := 2; k <= n; k++ {
			unique[kMap[k]] = true
		}
		for v := range unique {
			sum += v
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args:     []string{"12_000"},
			want:     "7587457",
			estimate: 1,
		},
		{
			args: []string{"12"},
			want: "61",
		},
		{
			args: []string{"6"},
			want: "30",
		},
	})
}

type n88 struct {
	remaining int
	sum       int
	f         int
}

type ctx88 struct {
	factors []int
	n       int
	kMap    map[int]int
	g       *generator.Prime
}

func (n *n88) Code(ctx *ctx88, dp bfs.DFSPath[*n88]) string {
	return fmt.Sprintf("%d,%d,%d", n.remaining, n.sum, dp.Len())
}

func (n *n88) Done(ctx *ctx88, dp bfs.DFSPath[*n88]) bool {
	if n.remaining != 1 {
		return false
	}
	if n.sum > ctx.n {
		return false
	}

	numberOfOnes := ctx.n - n.sum
	numberOfTerms := numberOfOnes + len(ctx.factors) - 1
	m := ctx.kMap
	if v, ok := m[numberOfTerms]; ok {
		if ctx.n < v {
			m[numberOfTerms] = ctx.n
		}
	} else {
		m[numberOfTerms] = ctx.n
	}
	return false
}

func (n *n88) OnPush(ctx *ctx88, dp bfs.DFSPath[*n88]) {
	ctx.factors = append(ctx.factors, n.f)
}

func (n *n88) OnPop(ctx *ctx88, dp bfs.DFSPath[*n88]) {
	ctx.factors = ctx.factors[:len(ctx.factors)-1]
}

func (n *n88) AdjacentStates(ctx *ctx88, dp bfs.DFSPath[*n88]) []*n88 {
	if n.remaining == 1 {
		return nil
	}
	if n.sum >= ctx.n {
		return nil
	}

	minFactor := 2
	if len(ctx.factors) > 0 {
		minFactor = ctx.factors[len(ctx.factors)-1]
	}
	var r []*n88
	for _, i := range ctx.g.Factors(n.remaining) {
		if i == 1 || i < minFactor {
			continue
		}
		r = append(r, &n88{n.remaining / i, n.sum + i, i})
	}
	return r
}
