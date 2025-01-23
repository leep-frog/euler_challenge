package p211

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P211() *ecmodels.Problem {
	return ecmodels.IntInputNode(211, func(o command.Output, n int) {

		ctx := &context{
			p:   generator.Primes(),
			sum: 0,
			max: n,
		}
		one := &node{
			k:                1,
			divisorSquareSum: 1,
			pIdx:             -1,
		}

		dfs(ctx, one)

		o.Stdoutln(ctx.sum)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"64000000"},
			Want:     "1922364685",
			Estimate: 25,
		},
	})
}

type node struct {
	k                int
	divisorSquareSum int
	pIdx             int
}

type context struct {
	p   *generator.Prime
	sum int
	max int
}

func dfs(ctx *context, n *node) {

	if maths.IsSquare(n.divisorSquareSum) {
		ctx.sum += n.k
	}

	for i := n.pIdx + 1; n.k*ctx.p.Nth(i) < ctx.max; i++ {
		f := ctx.p.Nth(i)
		coefs := 1 + f*f
		for cnt, curF := 1, f; n.k*curF < ctx.max; cnt, curF = cnt+1, curF*f {
			next := &node{
				k:                n.k * curF,
				divisorSquareSum: n.divisorSquareSum * coefs,
				pIdx:             i,
			}
			dfs(ctx, next)
			coefs += curF * curF * f * f
		}
	}
}
