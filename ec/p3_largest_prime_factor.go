package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

// TODO: move all of these to helper directory

func P3() *problem {
	return intInputNode(3, func(o command.Output, n int) {
		factors := generator.PrimeFactors(n, generator.Primes())

		max := 0
		for f := range factors {
			if f > max {
				max = f
			}
		}

		o.Stdoutln(max)
	})
}
