package p50

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P50() *ecmodels.Problem {
	return ecmodels.IntInputNode(50, func(o command.Output, n int) {
		best := maths.Largest[int, int]()
		primes := generator.Primes()
		for i := 0; primes.Nth(i) < n; i++ {
			pi := primes.Nth(i)
			sum := pi + primes.Nth(i+1)
			for j := 2; sum < n; j++ {
				if primes.Contains(sum) {
					best.IndexCheck(sum, j)
				}
				sum += primes.Nth(i + j)
			}
		}
		o.Stdoutln(best.BestIndex(), best.Best())
	}, []*ecmodels.Execution{
		{
			Args: []string{"1000000"},
			Want: "997651 543",
		},
		{
			Args: []string{"1000"},
			Want: "953 21",
		},
		{
			Args: []string{"100"},
			Want: "41 6",
		},
	})
}
