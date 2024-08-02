package p710

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P710() *ecmodels.Problem {
	return ecmodels.NoInputWithExampleNode(710, func(o command.Output, ex bool) {
		mod := 1_000_000

		// The sequence of this is:
		// 0 1 0 2 1 4 3 9 7 20 16 43 36
		// Splitting by odd/even, we get:
		// 0 0 1 3 7 16 36
		// 1 2 4 9 20 43
		// This ends up being the first and second level deltas of https://oeis.org/A000253
		//            1 2 4  9  20
		//           0 1 3 7  16 36
		// A000253: 0 0 1 4 11 27 63

		// a(n) = 2*a(n-1) - a(n-2) + a(n-3) + 2^(n-1).
		prev1 := 4
		prev2 := 1
		prev3 := 0

		for i := 5; ; i++ {
			var v int
			if i%2 == 1 {
				next := (2*prev1 - prev2 + prev3 + maths.PowMod(2, i/2, mod)) % mod
				v = prev2 - prev3
				prev1, prev2, prev3 = next, prev1, prev2
			} else {
				v = (prev1 - prev2) - (prev2 - prev3)
			}

			for v < 0 {
				v = v + mod
			}

			if ex && i == 42 {
				o.Stdoutln(v)
				return
			} else if v%mod == 0 {
				o.Stdoutln(i)
				return
			}
		}

	}, []*ecmodels.Execution{
		{
			Args: []string{"-x"},
			Want: "999923",
		},
		{
			Want:     "1275000",
			Estimate: 1,
		},
	})
}

func brute(n, min int, hasTwo bool, cur []int) *maths.Int {
	if n == 0 {
		if hasTwo {
			return combinatorics.PermutationCount(cur)
		}
		return maths.Zero()
	}

	cnt := maths.Zero()

	// First, put a number in the middle
	if hasTwo || n == 2 {
		cnt = cnt.Plus(combinatorics.PermutationCount(cur))
	}

	// Now iterate over possible palindromes to add
	for i := min; i*2 <= n; i++ {
		cnt = cnt.Plus(brute(n-i*2, i, hasTwo || i == 2, append(cur, i)))
	}
	return cnt
}
