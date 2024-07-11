package p48

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P48() *ecmodels.Problem {
	return ecmodels.IntInputNode(48, func(o command.Output, n int) {
		res := 0
		largest := 10_000_000_000
		_ = largest
		for i := 1; i <= n; i++ {
			prod := i
			for j := 1; j < i; j++ {
				prod = (prod * i) % largest
			}
			res = (res + prod) % largest
		}
		o.Stdoutln(res % largest)
	}, []*ecmodels.Execution{
		{
			Args: []string{"1000"},
			Want: "9110846700",
		},
		{
			Args: []string{"10"},
			Want: "405071317",
		},
	})
}
