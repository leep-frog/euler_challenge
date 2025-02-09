package p95

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P95() *ecmodels.Problem {
	return ecmodels.NoInputNode(95, func(o command.Output) {
		n := 1_000_000

		m := map[int]int{}
		for i := 2; i <= n; i++ {
			m[i] = bread.Sum(maths.Divisors(i)) - i
		}

		chainLens := map[int]int{}
		for k := range m {
			if _, ok := chainLens[k]; ok {
				continue
			}

			var chain []int
			pos := map[int]int{}
			for ; k < n && pos[k] == 0; k = m[k] {
				chain = append(chain, k)
				pos[k] = len(chain)
			}

			if k >= n {
				for _, j := range chain {
					chainLens[j] = -1
				}
			} else {
				mark := pos[k] - 1
				for i, v := range chain {
					if i < mark {
						chainLens[v] = -1
					} else {
						chainLens[v] = len(pos) - mark
					}
				}
			}
		}

		bestChain := 0
		smallestPart := n
		for k, v := range chainLens {
			if v > bestChain {
				bestChain = v
				smallestPart = k
			} else if v == bestChain && k < smallestPart {
				smallestPart = k
			}
		}
		o.Stdoutln(bestChain, smallestPart)
	}, &ecmodels.Execution{
		Want:     "28 14316",
		Estimate: 7,
	})
}
