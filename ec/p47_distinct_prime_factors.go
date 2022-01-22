package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P47() *problem {
	return intInputNode(47, func(o command.Output, n int) {
		p := generator.Primes()
		var row int
		for i := 1; ; i++ {
			if len(generator.PrimeFactors(i, p)) >= n {
				row++
				if row == n {
					o.Stdoutln(i - (n - 1))
					return
				}
			} else {
				row = 0
			}
		}
	})
}
