package p885

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

const mod = 1123455689

func P885() *ecmodels.Problem {
	return ecmodels.IntInputNode(885, func(o command.Output, n int) {
		o.Stdoutln(clever(n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "45",
		},
		{
			Args: []string{"5"},
			Want: "1543545675",
		},
		{
			Args: []string{"18"},
			Want: "827850196",
			// Estimate: 90, // My solution took 90s
		},
	})
}

// This is taken from the problem thread
func clever(n int) int {
	summation := maths.Zero()
	for i := 1; i <= 9; i++ {
		summation = summation.Plus(maths.BigPow(10+9*i, n))
	}
	summation = summation.DivInt(9)
	summation = summation.Minus(maths.BigPow(10, n))
	return summation.ModInt(mod)
}

// This is my OG solution (which would have been THIRD if started when initially released!!!)
func combos(rem, min int, cur []int) int {
	if rem == 0 {
		return (combinatorics.PermutationCount(cur).ModInt(mod) * maths.IntFromDigits(cur).ModInt(mod)) % mod
	}

	var m int
	for i := min; i <= 9; i++ {
		m = (m + combos(rem-1, i, append(cur, i))) % mod
	}
	return m
}
