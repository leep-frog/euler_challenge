package p2

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P2() *ecmodels.Problem {
	return ecmodels.IntInputNode(2, func(o command.Output, n int) {
		var sum int
		for iter, i := generator.Fibonaccis().Start(0); i < n; i = iter.Next() {
			if i%2 == 0 {
				sum += i
			}
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"4000000"},
			Want: "4613732",
		},
	})
}
