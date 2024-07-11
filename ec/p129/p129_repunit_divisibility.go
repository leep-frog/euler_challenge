package p129

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P129() *ecmodels.Problem {
	return ecmodels.IntInputNode(129, func(o command.Output, n int) {
		diffs := []int{2, 4, 2, 2}
		n = maths.Pow(10, n)
		// Noticed pattern for 10^x solutions:
		// 10 1000003
		// 100 1000003
		// 1000 1000003
		// 10000 1000003
		// 100000 1000003
		// 1000000 1000023

		for i, j := n+3, 1; ; i, j = i+diffs[j], (j+1)%len(diffs) {
			if RepunitSmallest(i) >= n {
				o.Stdoutln(i)
				return
			}
		}
	}, []*ecmodels.Execution{
		{
			Args: []string{"6"},
			Want: "1000023",
		},
		{
			Args: []string{"1"},
			Want: "17",
		},
	})
}

// Repunitable returns whether or not n can be a factor of a repunit (111...).
func Repunitable(n int) bool {
	return n%2 != 0 && n%5 != 0
}

// RepunitSmallest returns the length of the smallest repunit (111...)
// that has n as a factor. Calling functions should verify the input is
// Repunitable before calling this function
func RepunitSmallest(n int) int {
	if !Repunitable(n) {
		panic(fmt.Sprintf("repunit for n=%d requires GCD(n, 10) = 0", n))
	}
	// Build map from one digit to required multiplier
	mults := make([]int, 10, 10)
	for m := 1; m <= 9; m++ {
		prod := n * m
		digits := maths.Digits(prod)
		mults[digits[len(digits)-1]] = m * n
	}

	k := 1
	for init := mults[1] / 10; init != 0; {
		k++
		need := (11 - (init % 10)) % 10
		init = (init + mults[need]) / 10
	}
	return k
}
