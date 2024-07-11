package p12

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P12() *ecmodels.Problem {
	return ecmodels.IntInputNode(12, func(o command.Output, n int) {
		ts := generator.Triangulars()
		for idx, i := 0, ts.Nth(0); ; idx, i = idx+1, ts.Nth(idx+1) {
			var count int
			max := i / 2
			for j := 1; j < max; j++ {
				if i%j == 0 {
					max = (i / j) - 1
					if j*j == i {
						count++
					} else {
						count += 2
					}
					if count > n {
						o.Stdoutln(i)
						return
					}
				}
			}
		}
	}, []*ecmodels.Execution{
		{
			Args: []string{"5"},
			Want: "28",
		},
		{
			Args:     []string{"500"},
			Want:     "76576500",
			Estimate: 0.2,
		},
	})
}
