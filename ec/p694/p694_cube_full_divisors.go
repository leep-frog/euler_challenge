package p694

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

var (
	p = generator.Primes()
)

func P694() *ecmodels.Problem {
	return ecmodels.IntInputNode(694, func(o command.Output, n int) {
		o.Stdoutln(n + dp(0, n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"16"},
			Want: "19",
		},
		{
			Args: []string{"100"},
			Want: "126",
		},
		{
			Args: []string{"10000"},
			Want: "13344",
		},
		{
			Args: []string{"1000000000000000000"},
			Want: "1339784153569958487",
		},
	})
}

func dp(primeIdx, n int) int {
	if n == 0 {
		return 0
	}

	prime := p.Nth(primeIdx)
	if prime*prime*prime > n {
		return 0
	}

	// Continue when this prime is not present at all
	sum := dp(primeIdx+1, n)

	// Iterate over prime^3, prime^4, prime^5, ...
	for m := n / (prime * prime * prime); m > 0; m /= prime {
		sum += m
		sum += dp(primeIdx+1, m)
	}
	return sum
}
