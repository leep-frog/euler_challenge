package p407

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

var (
	p = generator.Primes()
)

func P407() *ecmodels.Problem {
	return ecmodels.IntInputNode(407, func(o command.Output, n int) {

		var bests []int
		for i := 0; i <= n; i++ {
			bests = append(bests, 1)
		}

		prevFactors := []int{1}
		for a := 2; a < n; a++ {
			// a^2 == a (mod k)
			// a^2 = a + x*k
			// a^2 - a = (x*k)
			// a*(a-1) = (x*k)
			curFactors := p.Factors(a)
			for i := len(curFactors) - 1; i >= 0; i-- {
				f1 := curFactors[i]
				for j := len(prevFactors) - 1; j >= 0; j-- {
					f2 := prevFactors[j]
					factor := f1 * f2
					if a >= factor {
						break
					}
					if factor <= n {
						bests[factor] = a
					} else {
					}
				}
			}
			prevFactors = curFactors
		}

		o.Stdoutln(bread.Sum(bests[2:]))
	}, []*ecmodels.Execution{
		{
			Args: []string{"10"},
			Want: "17",
		},
		{
			Args: []string{"1000"},
			Want: "314034",
		},
		{
			Args: []string{"10000000"},
			Want: "39782849136421",
		},
	})
}

func brute(n int) int {
	var sum int
	for j := 2; j <= n; j++ {
		best := 1
		for a := j - 1; a >= 2; a-- {
			if a*a%j == a {
				best = a
				break
			}
		}
		sum += best
	}
	return sum
}
