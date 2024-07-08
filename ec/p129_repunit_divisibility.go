package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P129() *problem {
	return intInputNode(129, func(o command.Output, n int) {
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
			if repunitSmallest(i) >= n {
				o.Stdoutln(i)
				return
			}
		}
	}, []*execution{
		{
			args: []string{"6"},
			want: "1000023",
		},
		{
			args: []string{"1"},
			want: "17",
		},
	})
}

// repunitable returns whether or not n can be a factor of a repunit (111...).
func repunitable(n int) bool {
	return n%2 != 0 && n%5 != 0
}

// repunitSmallest returns the length of the smallest repunit (111...)
// that has n as a factor. Calling functions should verify the input is
// repunitable before calling this function
func repunitSmallest(n int) int {
	if !repunitable(n) {
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
