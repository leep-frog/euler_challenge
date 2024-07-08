package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
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

func P172() *problem {
	return intInputNode(172, func(o command.Output, n int) {
		fmt.Println("START")
		o.Stdoutln(recursive172(n, 1, nil))
	}, []*execution{
		{
			args: []string{"1"},
			want: "9",
		},
		{
			args: []string{"2"},
			want: "90",
		},
		{
			args: []string{"3"},
			want: "900",
		},
		{
			args: []string{"4"},
			want: "8991",
		},
		{
			args: []string{"5"},
			want: "89586",
		},
		{
			args: []string{"6"},
			want: "888570",
		},
		{
			args: []string{"18"},
			want: "227485267000992000",
		},
	})
}
