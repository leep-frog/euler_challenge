package p187

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"golang.org/x/exp/slices"
)

func P187() *ecmodels.Problem {
	return ecmodels.IntInputNode(187, func(o command.Output, n int) {
		p := generator.Primes()

		var count int
		var relevantPrimes []int
		for i := 0; 2*p.Nth(i) < n; i++ {
			relevantPrimes = append(relevantPrimes, p.Nth(i))
		}

		for idx, rp := range relevantPrimes {

			var bruteCount int
			j := idx
			for ; rp*p.Nth(j) < n; j++ {
				bruteCount++
			}

			pos, at := slices.BinarySearch(relevantPrimes, n/rp)
			if at {
				pos++
			}
			if pos > idx {
				count += pos - idx
			}
		}

		o.Stdoutln(count)
	}, []*ecmodels.Execution{
		{
			Args: []string{"30"},
			Want: "10",
		},
		{
			Args:     []string{"100_000_000"},
			Want:     "17427258",
			Estimate: 14,
		},
	})
}
