package p57

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P57() *ecmodels.Problem {
	return ecmodels.NoInputNode(57, func(o command.Output) {
		num, den := maths.NewInt(3), maths.NewInt(2)
		var count int
		for i := 0; i < 1000; i++ {
			tmp := den
			den = den.Plus(num)
			num = tmp.Times(maths.NewInt(2)).Plus(num)
			if len(num.Digits()) > len(den.Digits()) {
				count++
			}
		}
		o.Stdoutln(count)
	}, &ecmodels.Execution{
		Want: "153",
	})
}
