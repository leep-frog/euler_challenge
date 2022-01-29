package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P50() *problem {
	return intInputNode(50, func(o command.Output, n int) {
		best := maths.Largest[int, int]()
		primes := generator.Primes()
		for i := 0; primes.Nth(i) < n; i++ {
			pi := primes.Nth(i)
			sum := pi + primes.Nth(i+1)
			for j := 2; sum < n; j++ {
				if generator.IsPrime(sum, primes) {
					best.IndexCheck(sum, j)
				}
				sum += primes.Nth(i + j)
			}
		}
		o.Stdoutln(best.BestIndex(), best.Best())
	})
}
