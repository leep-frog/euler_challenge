package p1

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P1() *ecmodels.Problem {
	return ecmodels.IntInputNode(1, func(o command.Output, n int) {
		var sum int
		for i := 1; i < n; i++ {
			if i%5 == 0 || i%3 == 0 {
				sum += i
			}
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			// Example
			Args: []string{"10"},
			Want: "23",
		},
		{
			Args: []string{"1000"},
			Want: "233168",
		},
	})
}
