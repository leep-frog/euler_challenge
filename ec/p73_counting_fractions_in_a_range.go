package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P73() *problem {
	return intInputNode(73, func(o command.Output, n int) {
		p := generator.Primes()
		var unique int
		for den := 4; den <= n; den++ {
			for num := den / 3; ; num++ {
				if num*2 >= den {
					break
				}
				if num*3 <= den {
					continue
				}
				if generator.Coprimes(num, den, p) {
					continue
				}
				unique++
			}
		}
		o.Stdoutln(unique)
	})
}
