package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P27() *problem {
	return intInputNode(27, func(o command.Output, n int) {
		p := generator.Primes()

		var max, maxI int
		for a := -n + 1; a < n; a++ {
			for b := -n; b <= n; b++ {
				// Try positive direction
				k := 0

				for ; generator.IsPrime(k*k+a*k+b, p); k++ {
				}
				if k > max {
					max = k
					maxI = a * b
				}

				// Try negative direction
				k = 0
				for ; generator.IsPrime(k*k+a*k+b, p); k-- {
				}
				if k > max {
					max = k
					maxI = a * b
				}
			}
		}
		o.Stdoutln(maxI)
	})
}
