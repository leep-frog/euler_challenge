package p193

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P193() *ecmodels.Problem {
	return ecmodels.IntInputNode(193, func(o command.Output, pow int) {

		n := maths.Pow(2, pow)

		p := generator.Primes()
		var sum int
		for i := 0; p.Nth(i)*p.Nth(i) <= n; i++ {
			sum += dp(p, n, i)
		}
		o.Stdoutln(n - sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"4"},
			Want: "11",
		},
		{
			Args:     []string{"50"},
			Want:     "684465067343069",
			Estimate: 10,
		},
	})
}

var (
	cache = map[string]int{}
)

func dp(p *generator.Prime, n, pIdx int) int {
	// 4 removes every 4th number
	// 9 removes every 9th number (but re-add every 4th combination)

	// When adding 2, f(n, 2): (n / 4)
	// When adding 3, f(n, 3): + (n / 9) - (n / 4,9)
	// When adding 5, f(n, 5): + (n / 25) - (n / 25,4) - (n / 25,9) + (n / 25,9,4)
	// When adding 7, f(n, 7): + (n / 49) - (n / 49,4) - (n / 49,9) - (n / 49,25) + (n / 49,4,9) + (n / 49,4,25) + (n / 49,9,25) - (n / 49,4,9,25)
	//
	// Therefore, f(n, x) = (n / x^2) - ( f(n/x^2, x-1) + f(n/x^2, x - 2) + ... + f(n/x^2, 3) + f(n/x^2, 2) )
	//
	// And solution is sum from x = 2,3,5,7,... of f(n, x)   (This part is done in the top-level method)

	if n == 0 {
		return 0
	}

	nextN := n / (p.Nth(pIdx) * p.Nth(pIdx))

	sum := nextN

	code := fmt.Sprintf("%d-%d", pIdx, nextN)
	if v, ok := cache[code]; ok {
		return v
	}

	for i := 0; i < pIdx; i++ {
		if v := dp(p, nextN, i); v == 0 {
			break
		} else {
			sum -= v
		}
	}

	cache[code] = sum
	return sum
}
