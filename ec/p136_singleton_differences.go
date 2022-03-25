package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P136() *problem {
	return intInputNode(136, func(o command.Output, max int) {
		g := generator.Primes()
		var count int
		for n := 2; n < max; n++ {
			if diophantineDifferenceExactCount(n, 1, g) {
				count++
			}
		}
		o.Stdoutln(count)
	})
}
