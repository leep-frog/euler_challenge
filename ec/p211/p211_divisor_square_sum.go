package p211

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P211() *ecmodels.Problem {
	return ecmodels.IntInputNode(211, func(o command.Output, n int) {
		p := generator.Primes()

		var sum int
		for i := 1; i < n; i++ {
			var cur int
			pfn := p.PrimeFactoredNumberFast(i)
			for _, d := range pfn.PrimeDivisors(p) {
				cur += d * d
			}
			if maths.IsSquare(cur) {
				fmt.Println("YUP", i)
				sum += i
			}
		}

		ctx := &context{p, 0, n}
		one := &node{
			k:                1,
			divisorSquareSum: 1,
			pIdx:             -1,
		}
		bfs.ContextSearch[string, *context, *node](ctx, []*node{one})

		o.Stdoutln(n, ctx.sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"64000000"},
			Want: "1922364685",
		},
	})
}

// func o2(p *generator.Prime, k int) int {
// ;p.
// }

// 42 = 2 * 21 = 2 * 3 * 7
// 1 2 3 6 7 12 18

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

func (n *node) Done(ctx *context) bool {
	if maths.IsSquare(n.divisorSquareSum) {
		ctx.sum += n.k
	}
	return false
}

func (n *node) Code(ctx *context) string {
	return fmt.Sprintf("%d", n.k)
}

func (n *node) AdjacentStates(ctx *context) []*node {
	var ns []*node

	for i := n.pIdx + 1; n.k*ctx.p.Nth(i) < ctx.max; i++ {
		f := ctx.p.Nth(i)
		coefs := 1 + f
		for cnt, curF := 1, f; n.k*f < ctx.max; cnt, curF = cnt+1, curF*f {
			ns = append(ns, &node{
				k:                n.k * f,
				divisorSquareSum: n.divisorSquareSum * coefs,
				pIdx:             i,
			})
			coefs += curF
		}
	}
	return ns
}

// 1 + 42 + 246 + 287 + 728 + 1434 + 1673 + 1880 + 4264 + 6237 + 9799 + 9855 + 18330 + 21352 + 21385 + 24856 + 36531 + 39990 + 46655 + 57270 + 66815 + 92664 + 125255 + 156570 + 182665 + 208182 + 212949 + 242879 + 273265 + 380511 + 391345 + 411558 + 539560 + 627215 + 693160 + 730145 + 741096 + 773224 + 814463 + 931722 + 992680 + 1069895 + 1087009 + 1143477 + 1166399 + 1422577 + 1592935 + 1815073 + 2281255 + 2544697 + 2713880 + 2722005 + 2828385 + 3054232 + 3132935 + 3145240 + 3188809 + 3508456 + 4026280 + 4647985 + 4730879 + 5024488 + 5054015 + 5143945 + 5260710 + 5938515 + 6128024 + 6236705 + 6366767 + 6956927 + 6996904 + 7133672 + 7174440 + 7538934 + 7736646 + 7818776 + 8292583 + 8429967 + 8504595 + 8795423 + 9026087 + 9963071 + 11477130 + 11538505 + 11725560 + 12158135 + 12939480 + 12948776 + 13495720 + 13592118 + 13736408 + 15203889 + 15857471 + 16149848 + 16436490 + 16487415 + 16909849 + 18391401 + 18422120 + 20549528 + 20813976 + 20871649 + 21251412 + 22713455 + 23250645 + 23630711 + 24738935 + 26338473 + 26343030 + 26594568 + 28113048 + 29429144 + 29778762 + 29973414 + 30666090 + 30915027 + 34207446 + 34741889 + 34968983 + 35721896 + 37113593 + 37343065 + 38598255 + 39256230 + 42021720 + 44935590 + 45795688 + 45798935 + 48988758 + 49375521 + 51516049 + 51912289 + 52867081 + 56215914 + 59748234 + 61116363 + 62158134 + 63286535
