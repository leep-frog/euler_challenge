package p77

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P77() *ecmodels.Problem {
	return ecmodels.NoInputNode(77, func(o command.Output) {
		p := generator.Primes()
		sumMap := map[int]int{}
		dfs77(0, 0, 0, sumMap, p)
		best := maths.Largest[int, int]()
		for k, v := range sumMap {
			best.IndexCheck(k, v)
		}
		o.Stdoutln(best.BestIndex(), best.Best())
	}, &ecmodels.Execution{
		Want: "71 5006",
	})
}

func dfs77(sum, idx, depth int, m map[int]int, p *generator.Prime) {
	if depth >= 2 {
		m[sum]++
	}
	// Just binary searched for max until the best was over 5000
	for i := idx; p.Nth(i)+sum <= 71; i++ {
		dfs77(sum+p.Nth(i), i, depth+1, m, p)
	}
}
