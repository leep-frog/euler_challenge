package p135

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P135() *ecmodels.Problem {
	return ecmodels.IntInputNode(135, func(o command.Output, max int) {
		g := generator.Primes()
		var count int
		for n := 2; n < max; n++ {
			if DiophantineDifferenceExactCount(n, 10, g) {
				count++
			}
		}
		o.Stdoutln(count)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"1000000"},
			Want:     "4989",
			Estimate: 1,
		},
		{
			Args: []string{"1156"},
			Want: "1",
		},
	})
}

// DiophantineDifferenceCount returns true if the number of solutions to
// x^2 - y^2 - z^2 = n
// is exactly count
func DiophantineDifferenceExactCount(n, count int, primes *generator.Prime) bool {
	// Noticed in all solutions have y as a factor of n
	var solutions int
	for _, f := range primes.Factors(n) {
		if f == 1 {
			continue
		}
		// (f+k)^2 - f^2 - (f-k)^2 = n
		// f^2 + 2fk + k^2 - f^2 - f^2 + 2fk - k^2 = n
		// 4fk - f^2 = n
		// 4fk = n + f^2
		// k = (n+f^2)/4f
		k := (n + f*f) / (4 * f)
		if k > 0 && k < f && k*4*f == n+f*f {
			solutions++
		}
		if solutions > count {
			return false
		}
	}
	return solutions == count
}
