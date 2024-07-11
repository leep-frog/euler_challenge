package p28

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P28() *ecmodels.Problem {
	return ecmodels.IntInputNode(28, func(o command.Output, n int) {
		start := 3
		sum := 1
		for i := 0; i < (n-1)/2; i++ {
			offset := (i + 1) * 2
			sum += 4*start + 6*offset
			start += offset*4 + 2
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"1001"},
			Want: "669171001",
		},
		{
			Args: []string{"5"},
			Want: "101",
		},
	})
}
