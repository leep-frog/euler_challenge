package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P135() *problem {
	return intInputNode(135, func(o command.Output, max int) {
		g := generator.Primes()
		var count int
		for n := 2; n < max; n++ {
			if diophantineDifferenceExactCount(n, 10, g) {
				count++
			}
		}
		o.Stdoutln(count)
	})
}

// diophantineDifferenceCount returns true if the number of solutions to
// x^2 - y^2 - z^2 = n
// is exactly count
func diophantineDifferenceExactCount(n, count int, primes *generator.Generator[int]) bool {
	// Noticed in all solutions have y as a factor of n
	var solutions int
	for _, f := range generator.Factors(n, primes) {
		if f == 1 {
			continue
		}
		//fmt.Println(n, f)
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
