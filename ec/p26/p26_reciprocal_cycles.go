package p26

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func cycleLen(num int) int {
	rem := 1
	remMap := map[int]int{}

	for pos := 0; ; pos++ {
		rem = (rem % num) * 10
		if rem == 0 {
			return 0
		}
		if v, ok := remMap[rem]; ok {
			return pos - v
		}
		remMap[rem] = pos
	}
}

func P26() *ecmodels.Problem {
	return ecmodels.IntInputNode(26, func(o command.Output, n int) {
		var max, maxI int
		for i := 1; i < n; i++ {
			if v := cycleLen(i); v > max {
				max = v
				maxI = i
			}
		}
		o.Stdoutln(maxI)
	}, []*ecmodels.Execution{
		{
			Args: []string{"1000"},
			Want: "983",
		},
		{
			Args: []string{"10"},
			Want: "7",
		},
	})
}
