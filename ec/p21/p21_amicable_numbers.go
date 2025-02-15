package p21

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P21() *ecmodels.Problem {
	return ecmodels.IntInputNode(21, func(o command.Output, n int) {

		sumMap := map[int]int{}
		for i := 1; i <= n; i++ {
			count := 0
			for _, div := range maths.Divisors(i) {
				if div != i {
					count += div
				}
			}
			sumMap[i] = count
		}

		var sum int
		for k, v := range sumMap {
			if sumMap[v] == k && v != k {
				sum += k + v
			}
		}

		// Divide by 2 since each pair is counted twice ((k, v) and (v, k)).
		o.Stdoutln(sum / 2)
	}, []*ecmodels.Execution{
		{
			Args: []string{"10000"},
			Want: "31626",
		},
	})
}
