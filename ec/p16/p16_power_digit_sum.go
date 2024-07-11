package p16

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P16() *ecmodels.Problem {
	return ecmodels.IntInputNode(16, func(o command.Output, ni int) {
		n := maths.NewInt(ni)

		two := maths.NewInt(2)
		pow := maths.NewInt(1)
		for n.GT(maths.Zero()) {
			pow = pow.Times(two)
			n.MM()
		}

		o.Stdoutln(pow.DigitSum())
	}, []*ecmodels.Execution{
		{
			Args: []string{"10"},
			Want: "7",
		},
		{
			Args: []string{"1000"},
			Want: "1366",
		},
	})
}
