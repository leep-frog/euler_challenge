package p64

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P64() *ecmodels.Problem {
	return ecmodels.IntInputNode(64, func(o command.Output, n int) {
		var count int
		for k := 2; k <= n; k++ {
			_, period := maths.SquareRootPeriod(k)
			if len(period)%2 == 1 {
				count++
			}
		}
		o.Stdoutln(count)
	}, []*ecmodels.Execution{
		{
			Args: []string{"10000"},
			Want: "1322",
		},
		{
			Args: []string{"13"},
			Want: "4",
		},
	})
}
