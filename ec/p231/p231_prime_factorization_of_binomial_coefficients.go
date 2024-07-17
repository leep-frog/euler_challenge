package p231

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P231() *ecmodels.Problem {
	return ecmodels.NoInputWithExampleNode(231, func(o command.Output, example bool) {
		p := generator.Primes()

		a, b := 10, 3
		if !example {
			a, b = 20_000_000, 15_000_000
		}

		// 10 choose 7 === 10 choose 3
		b = maths.Min(b, a-b)

		m := map[int]int{}

		// a choose b = a! / ( (a-b)! * b! )
		// Assume b < a/2, then we get
		// a choose b = a * (a-1) * ... * (a - b + 1) / (b * (b - 1) * ... 2 * 1)

		// Increment factors for [ a * (a-1) * ... * (a - b + 1) ]
		for i := a - b + 1; i <= a; i++ {
			for f, cnt := range p.PrimeFactors(i) {
				m[f] += cnt
			}
		}

		// Decrement factors for (b * (b - 1) * ... 2 * 1)
		for i := 1; i <= b; i++ {
			for f, cnt := range p.PrimeFactors(i) {
				m[f] -= cnt
			}
		}

		var sum int
		for f, cnt := range m {
			sum += f * cnt
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"-x"},
			Want: "14",
		},
		{
			Want: "7526965179680",
		},
	})
}
