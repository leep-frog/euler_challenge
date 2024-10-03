package p429

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

const mod = 1_000_000_009

func P429() *ecmodels.Problem {
	return ecmodels.IntInputNode(429, func(o command.Output, n int) {

		p := generator.Primes()

		sum := 1
		for i := 0; p.Nth(i) <= n; i++ {

			prime := p.Nth(i)

			var cnt int
			for k := n / prime; k > 0; k = k / prime {
				cnt += k
			}

			left, right := sum, (sum*maths.PowMod(prime, 2*cnt, mod))%mod
			sum = (left + right) % mod
		}

		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"4"},
			Want: "650",
		},
		{
			Args:     []string{"100000000"},
			Want:     "98792821",
			Estimate: 18,
		},
	})
}
