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
			ctx := &context88{i, kMap, g}
			bfs.AnyPath([]*node88{{nil, i, 0}}, ctx)
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
	})
}

type node88 struct {
	factors []int
	remaining int
	sum int
}

type context88 struct {
	n int
	kMap map[int]int
	g *generator.Generator[int]
}

func (n *node88) Code(*bfs.Context[*context88, *node88]) string {
	//c := fmt.Sprintf("%d: %v", n.remaining, n.factors)
	//fmt.Println(c)
	return fmt.Sprintf("%d: %v", n.remaining, n.factors)
}

func (n *node88) Done(ctx *bfs.Context[*context88, *node88]) bool {
	if n.remaining != 1 {
		return false
	}
	if n.sum > ctx.GlobalContext.n {
		return false
	}
	numberOfOnes := ctx.GlobalContext.n - n.sum
	numberOfTerms := numberOfOnes + len(n.factors)
	m := ctx.GlobalContext.kMap
	if v, ok := m[numberOfTerms]; ok {
		if ctx.GlobalContext.n < v {
			m[numberOfTerms] = ctx.GlobalContext.n
		}
	} else {
		m[numberOfTerms] = ctx.GlobalContext.n
	}
	return false
}

func (n *node88) AdjacentStates(ctx *bfs.Context[*context88, *node88]) []*node88 {
	if n.remaining == 1 {
		return nil
	}
	if n.sum >= ctx.GlobalContext.n {
		return nil
	}

	var r []*node88
	for _, i := range generator.Factors(n.remaining, ctx.GlobalContext.g) {
		if i == 1 {
			continue
		}
		r = append(r, &node88{append(n.factors, i), n.remaining/i, n.sum + i})	
	}
	return r
}
