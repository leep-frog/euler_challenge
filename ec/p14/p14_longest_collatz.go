package p14

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P14() *ecmodels.Problem {
	return ecmodels.IntInputNode(14, func(o command.Output, n int) {
		found := map[int]int{}
		var max, maxI int
		for i := 2; i < n; i++ {
			count := 1
			for j := i; j != 1; {
				if j%2 == 0 {
					j /= 2
				} else {
					j = 3*j + 1
				}

				if v, ok := found[j]; ok {
					count += v
					break
				}
				count++
			}
			found[i] = count
			if count > max {
				max = count
				maxI = i
			}
		}
		o.Stdoutln(maxI)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"1000000"},
			Want:     "837799",
			Estimate: 0.4,
		},
	})
}
