package p6

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P6() *ecmodels.Problem {
	return ecmodels.IntInputNode(6, func(o command.Output, n int) {
		var squareSums, sumSquares int
		for i := 1; i <= n; i++ {
			sumSquares += i * i
			squareSums += i
		}
		squareSums *= squareSums
		o.Stdoutln(squareSums - sumSquares)
	}, []*ecmodels.Execution{
		{
			Args: []string{"10"},
			Want: "2640",
		},
		{
			Args: []string{"100"},
			Want: "25164150",
		},
	})
}
