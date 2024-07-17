package p172

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func recursive172(rem, min int, cur []int) *maths.Int {
	// No options if we're out of numbers.
	if len(cur) > 10 {
		return maths.Zero()
	}

	// Still have to add more numbers.
	if rem != 0 {
		sum := maths.Zero()
		for i := min; i <= 3 && i <= rem; i++ {
			sum = sum.Plus(recursive172(rem-i, i, append(cur, i)))
		}
		return sum
	}

	// Get the number of ways the parts can be arranged.
	var parts []int
	counts := make([]int, 3, 3)
	for i, cnt := range cur {
		counts[cnt-1]++
		for j := 0; j < cnt; j++ {
			parts = append(parts, i)
		}
	}
	pc := combinatorics.PermutationCount(parts)

	// We want to consider swapping numbers identical. Numbers that
	// appear the same number of times can be rearranged in all possible ways.
	div := 1
	for _, cnt := range counts {
		div *= maths.FactorialI(cnt)
	}

	// Numbers can appear in any order, but 0 can't be first (hence why we have 9 numbers to choose from).
	mul := []int{
		9, 9, 8, 7, 6, 5, 4, 3, 2, 1,
	}
	base := 1
	for i := 0; i < len(cur); i++ {
		base *= mul[i]
	}

	// Number of ways we can organize the parts,
	// divided by identical patterns (AABCDC == CCBADA == AADCBC)
	// times all possible combinations of numbers.
	return pc.DivInt(div).TimesInt(base)
}

func P172() *ecmodels.Problem {
	return ecmodels.IntInputNode(172, func(o command.Output, n int) {
		o.Stdoutln(recursive172(n, 1, nil))
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "9",
		},
		{
			Args: []string{"2"},
			Want: "90",
		},
		{
			Args: []string{"3"},
			Want: "900",
		},
		{
			Args: []string{"4"},
			Want: "8991",
		},
		{
			Args: []string{"5"},
			Want: "89586",
		},
		{
			Args: []string{"6"},
			Want: "888570",
		},
		{
			Args: []string{"18"},
			Want: "227485267000992000",
		},
	})
}
