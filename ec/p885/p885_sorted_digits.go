package p885

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P885() *ecmodels.Problem {
	return ecmodels.IntInputNode(885, func(o command.Output, n int) {

		for i := 1; i <= 18; i++ {
			fmt.Println(i, combos(i, 0, nil))
		}
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "",
		},
		{
			Args: []string{"2"},
			Want: "",
		},
	})
}

func combos(rem, min int, cur []int) *maths.Int {
	if rem == 0 {
		return combinatorics.PermutationCount(cur).Times(maths.IntFromDigits(cur))
	}

	m := maths.Zero()
	for i := min; i <= 9; i++ {
		m = m.Plus(combos(rem-1, i, append(cur, i)))
	}

	return m
}
