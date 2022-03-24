package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P133() *problem {
	return intInputNode(133, func(o command.Output, n int) {
		g := generator.Primes()
		var sum int
		for i, pi := 0, g.Nth(0); pi < n; i, pi = i+1, g.Nth(i+1) {
			if !repunitable(pi) {
				sum += pi
				continue
			}
			factors := generator.PrimeFactors(repunitSmallest(pi), g)

			for k := range factors {
				if k != 2 && k != 5 {
					sum += pi
					break
				}
			}
		}
		o.Stdoutln(sum)
	})
}
