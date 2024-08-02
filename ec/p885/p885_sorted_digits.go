package p885

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

const mod = 1123455689

func P885() *ecmodels.Problem {
	return ecmodels.IntInputNode(885, func(o command.Output, n int) {

		for i := 1; i <= 18; i++ {
			fmt.Println(i, combos(i, 0, nil))
		}
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
			Args:     []string{"18"},
			Want:     "827850196",
			Estimate: 90,
		},
	})
}

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
