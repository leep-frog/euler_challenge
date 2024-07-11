package p3

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P3() *ecmodels.Problem {
	return ecmodels.IntInputNode(3, func(o command.Output, n int) {
		factors := generator.Primes().PrimeFactors(n)

		max := 0
		for f := range factors {
			if f > max {
				max = f
			}
		}

		o.Stdoutln(max)
	}, []*ecmodels.Execution{
		{
			Args: []string{"13195"},
			Want: "29",
		},
		{
			Args: []string{"600851475143"},
			Want: "6857",
		},
	})
}
