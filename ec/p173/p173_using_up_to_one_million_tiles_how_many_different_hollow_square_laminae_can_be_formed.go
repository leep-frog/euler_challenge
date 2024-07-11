package p173

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P173() *ecmodels.Problem {
	return ecmodels.IntInputNode(173, func(o command.Output, n int) {
		// All squares are divisble by four, so sum will always be
		// 4*(a + (a + 1) + ...) <= n/4
		// (a + (a + 1) + ...) <= n/4
		n = n / 4

		var sum int
		for _, base := range []int{2, 3} {
			// base + (base + 2) + (base + 4) + (...) < n
			// k*base + (2 + 4 + 6 + ... + 2*(k-1))
			// k*base + 2(1 + 2 + 3 + ... + (k - 1)) < n
			// k*base + 2((k-1)*k)/2 < n
			// k*base + ((k-1)*k) < n
			// k*base + k^2 - k < n
			// k^2 + k(base - 1) < n
			// k(k + base - 1) < n
			k := 1
			for k*(k+base-1) <= n {
				k++
			}
			k--

			for ; base <= n; base += 2 {
				// Can also solve k dynamically via quadratic:
				// [k(k + base - 1) - n < 0]
				// But k is strictly decreasing so solution is linear either way
				for k*(k+base-1) > n {
					k--
				}
				sum += k
			}
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"100"},
			Want: "41",
		},
		{
			Args: []string{"1_000_000"},
			Want: "1572729",
		},
	})
}
