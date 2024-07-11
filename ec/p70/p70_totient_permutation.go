package p70

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P70() *ecmodels.Problem {
	return ecmodels.IntInputNode(70, func(o command.Output, n int) {
		p := generator.Primes()

		// Probably two primes together
		best := maths.Smallest[int, float64]()
		for i := 0; p.Nth(i) < n; i++ {
			pi := p.Nth(i)
			for j := i + 1; pi*p.Nth(j) < n; j++ {
				pj := p.Nth(j)
				k := pi * pj
				phi := k - (k/pi + k/pj) + 1
				if combinatorics.Anagram(k, phi) {
					best.IndexCheck(k, float64(k)/float64(phi))
				}
			}
		}
		o.Stdoutln(best.BestIndex())
		return
	}, []*ecmodels.Execution{
		{
			Args:     []string{"10000000"},
			Want:     "8319823",
			Estimate: 2,
		},
	})
}
