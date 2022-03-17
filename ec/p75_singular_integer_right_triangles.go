package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P75() *problem {
	return intInputNode(75, func(o command.Output, L int) {
		// https://en.wikipedia.org/wiki/Pythagorean_triple
		// a = m^2 - n^2
		// b = 2mn
		// c = m^2 + n^2
		// L = 2m^2 + 2mn
		counts := map[int]int{}
		g := generator.Primes()
		for m := 2; 2*m*m+2*m <= L; m++ {
			for n := 1; n < m; n++ {
				if n%2 == 1 && m%2 == 1 {
					continue
				}
				if n > 1 && generator.Coprimes(m, n, g) {
					continue
				}
				a := m*m - n*n
				b := 2 * m * n
				c := m*m + n*n
				perimeter := a + b + c
				for l := perimeter; l <= L; l += perimeter {
					counts[l]++
				}
			}
		}
		var count int
		for _, v := range counts {
			if v == 1 {
				count++
			}
		}
		o.Stdoutln(count)
	})
}
