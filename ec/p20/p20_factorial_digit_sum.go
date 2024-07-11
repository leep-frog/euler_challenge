package p20

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P20() *ecmodels.Problem {
	return ecmodels.IntInputNode(20, func(o command.Output, n int) {
		o.Stdoutln(maths.Factorial(n).DigitSum())
	}, []*ecmodels.Execution{
		{
			Args: []string{"100"},
			Want: "648",
		},
		{
			Args: []string{"10"},
			Want: "27",
		},
	})
}
