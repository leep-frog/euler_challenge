package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P77() *problem {
	return noInputNode(77, func(o command.Output) {
		p := generator.Primes()
		sumMap := map[int]int{}
		dfs77(0, 0, 0, sumMap, p)
		best := maths.Largest[int]()
		for k, v := range sumMap {
			best.IndexCheck(k, v)
		}
		o.Stdoutln(best.BestIndex(), best.Best())
	})
}

func dfs77(sum, idx, depth int, m map[int]int, p *generator.Generator[int]) {
	if depth >= 2 {
		m[sum]++
	}
	// Just binary searched for max until the best was over 5000
	for i := idx; p.Nth(i)+sum <= 71; i++ {
		dfs77(sum+p.Nth(i), i, depth+1, m, p)
	}
}
