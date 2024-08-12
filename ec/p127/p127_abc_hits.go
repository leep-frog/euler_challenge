package p127

import (
	"strconv"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/ec/p124"
	"github.com/leep-frog/euler_challenge/generator"
)

func P127() *ecmodels.Problem {
	return ecmodels.IntInputNode(127, func(o command.Output, n int) {
		g := generator.Primes()
		var sum int
		for c := 1; c < n; c++ {
			if g.Contains(c) {
				continue
			}
			cRad := p124.CalcRadical(c, g)

			var nonOverlappingPrimes []int
			for i := 0; g.Nth(i)*cRad < c; i++ {
				pi := g.Nth(i)
				if c%pi != 0 {
					nonOverlappingPrimes = append(nonOverlappingPrimes, pi)
				}
			}

			if len(nonOverlappingPrimes) == 0 {
				continue
			}
			ctx := &context127{g, c, cRad, nonOverlappingPrimes, 0}
			initStates := []*node127{{1, 0}}
			bfs.ContextualDFS(initStates, ctx)
			sum += ctx.sum
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"120000"},
			Want:     "18407904",
			Estimate: 10,
		},
		{
			Args: []string{"1000"},
			Want: "12523",
		},
	})
}

type node127 struct {
	a         int
	factorIdx int
}

type context127 struct {
	g              *generator.Prime
	c              int
	cRad           int
	allowedFactors []int
	sum            int
}

func (n *node127) Code(*context127, bfs.DFSPath[*node127]) string {
	return strconv.Itoa(n.a)
}

func (n *node127) Done(ctx *context127, dp bfs.DFSPath[*node127]) bool {
	b := ctx.c - n.a
	if b <= n.a {
		return false
	}
	aRad := p124.CalcRadical(n.a, ctx.g)
	bRad := p124.CalcRadical(b, ctx.g)
	if aRad*bRad*ctx.cRad >= ctx.c {
		return false
	}
	// Don't need to check (a, c) since that is guaranteed based on AdjacentStates
	if !ctx.g.Coprimes(n.a, ctx.c) || !ctx.g.Coprimes(b, ctx.c) {
		return false
	}
	ctx.sum += ctx.c
	return false
}

func (n *node127) AdjacentStates(ctx *context127, path bfs.DFSPath[*node127]) []*node127 {
	var r []*node127
	for i := n.factorIdx; i < len(ctx.allowedFactors) && n.a*ctx.allowedFactors[i] < ctx.c/2; i++ {
		pi := ctx.allowedFactors[i]
		r = append(r, &node127{n.a * pi, i})
	}
	return r
}
