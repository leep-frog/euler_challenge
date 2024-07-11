package p601

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P601() *ecmodels.Problem {
	return ecmodels.IntsInputNode(601, 1, 1, func(o command.Output, ns []int) {
		primes := generator.Primes()

		// define streak(n) = k as the smallest positive integer k such that n+k is not divisible by k+1.
		// E.g:
		// 13 is divisible by 1
		// 14 is divisible by 2
		// 15 is divisible by 3
		// 16 is divisible by 4
		// 17 is NOT divisible by 5
		// So streak(13) = 4.

		// This just means that (n-1) is divisble by 1, 2, ..., k
		// 12 is divisble by 1, 2, 3, 4, but not 5
		// Therefore, we are just looking for all numbers that are divisble by 1, 2, ..., k, but not by k+1
		p := func(s, k int) int {

			// Get the prime factor count required to be divisble by all numbers <= s.
			primeFactor := map[int]int{}
			for i := 2; i <= s; i++ {
				for f, cnt := range primes.PrimeFactors(i) {
					primeFactor[f] = maths.Max(primeFactor[f], cnt)
				}
			}

			// Create the smallest number that is divisible by s.
			start := 1
			for f, pow := range primeFactor {
				start *= maths.Pow(f, pow)
			}

			// Count up to k in increments of start
			// This is for numbers (1 < n < k) so subtract 1 from k
			count := (k - 1) / start

			// P(n, k) is for numbers (1 < n < k) so don't include 1 if starting from that.
			if start == 1 {
				count--
			}

			// On a linear basis, we will hit a number that is also divisble by (s+1),
			// so we need to remove those.
			removeDivisor := 1
			for f, cnt := range primes.PrimeFactors(s + 1) {
				for i := primeFactor[f]; i < cnt; i++ {
					removeDivisor *= f
				}
			}
			remove := count / removeDivisor

			return count - remove
		}

		if len(ns) == 1 {
			var sum int
			for i := 1; i <= ns[0]; i++ {
				sum += p(i, maths.Pow(4, i))
			}
			o.Stdoutln(sum)
		} else {
			o.Stdoutln(p(ns[0], ns[1]))
		}
	}, []*ecmodels.Execution{
		{
			Args: []string{"3", "14"},
			Want: "1",
		},
		{
			Args: []string{"6", "1_000_000"},
			Want: "14286",
		},
		{
			Args: []string{"31"},
			Want: "1617243",
		},
	})
}
