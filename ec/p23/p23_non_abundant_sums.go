package p23

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P23() *ecmodels.Problem {
	return ecmodels.IntInputNode(23, func(o command.Output, n int) {
		var abundant []int
		for i := 1; i <= n; i++ {
			divs := maths.Divisors(i)
			var sum int
			for _, d := range divs {
				if d != i {
					sum += d
				}
			}
			if sum > i {
				abundant = append(abundant, i)
			}
		}

		isSum := map[int]bool{}
		for i, ai := range abundant {
			for j := i; j < len(abundant); j++ {
				aj := abundant[j]
				isSum[ai+aj] = true
			}
		}

		var sum int
		for i := 1; i <= n; i++ {
			if !isSum[i] {
				sum += i
			}
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"28123"},
			Want:     "4179871",
			Estimate: 0.6,
		},
	})
}
