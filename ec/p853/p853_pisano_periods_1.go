package p853

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P853() *ecmodels.Problem {
	return ecmodels.NoInputWithExampleNode(853, func(o command.Output, ex bool) {

		v, max := 120, 1_000_000_000
		if ex {
			v, max = 18, 50
		}

		var sum int
		for i := 3; i < max; i++ {
			if pisano(i, v) {
				sum += i
				fmt.Println(i)
			}
		}
		fmt.Println(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"-x"},
			Want: "57",
		},
		{
			Want:     "44511058204",
			Estimate: 300,
		},
	})
}

func pisano(v, want int) bool {
	cur, prev := 1, 1

	cnt := 1
	for ; !(cur == 1 && prev == 0); cnt++ {
		cur, prev = (cur+prev)%v, cur
		if cnt > want {
			return false
		}
	}

	return cnt == want
}
