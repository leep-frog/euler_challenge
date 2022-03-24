package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P5() *problem {
	return intInputNode(5, func(o command.Output, n int) {
		primer := generator.Primes()
		primes := map[int]int{}
		for i := 2; i < n; i++ {
			for p, cnt := range generator.PrimeFactors(i, primer) {
				primes[p] = maths.Max(cnt, primes[p])
			}
		}
		product := 1
		for p, cnt := range primes {
			for i := 0; i < cnt; i++ {
				product *= p
			}
		}
		o.Stdoutln(product)
	})
}
