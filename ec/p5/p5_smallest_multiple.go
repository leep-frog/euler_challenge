package p5

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P5() *ecmodels.Problem {
	return ecmodels.IntInputNode(5, func(o command.Output, n int) {
		primer := generator.Primes()
		primes := map[int]int{}
		for i := 2; i < n; i++ {
			for p, cnt := range primer.PrimeFactors(i) {
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
	}, []*ecmodels.Execution{
		{
			Args: []string{"10"},
			Want: "2520",
		},
		{
			Args: []string{"20"},
			Want: "232792560",
		},
	})
}
