package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P23() *problem {
	return intInputNode(23, func(o command.Output, n int) {
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
	}, []*execution{
		{
			args:     []string{"28123"},
			want:     "4179871",
			estimate: 0.6,
		},
	})
}
