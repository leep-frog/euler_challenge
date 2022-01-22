package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P69() *problem {
	return intInputNode(69, func(o command.Output, n int) {
		p := generator.Primes()
		prod := 1
		for ; prod < n; prod *= p.Next() {
		}
		o.Stdoutln(prod / p.Last())
	})
}
