package p72

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P72() *ecmodels.Problem {
	return ecmodels.IntInputNode(72, func(o command.Output, n int) {
		p := generator.Primes()
		got := map[int]int{}
		count := 0

		for i := 2; i <= n; i++ {
			if p.Contains(i) {
				got[i] = i - 1
			} else {
				k := 1
				for num, cnt := range p.PrimeFactors(i) {
					k *= got[num] * maths.Pow(num, cnt-1)
				}
				got[i] = k
			}
			count += got[i]
		}
		o.Stdoutln(count)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"1000000"},
			Want:     "303963552391",
			Estimate: 2,
		},
		{
			Args: []string{"8"},
			Want: "21",
		},
	})
}
