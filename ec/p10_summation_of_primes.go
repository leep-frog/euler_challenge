package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P10() *problem {
	return intInputNode(10, func(o command.Output, n int) {
		p := generator.Primes()

		var sum int
		for pn := p.Next(); pn < n; pn = p.Next() {
			sum += pn
		}
		o.Stdoutln(sum)
	})
}
