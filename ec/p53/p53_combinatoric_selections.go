package p53

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P53() *ecmodels.Problem {
	return ecmodels.NoInputNode(53, func(o command.Output) {
		var count int
		mill := maths.NewInt(1_000_000)
		for n := 23; n <= 100; n++ {
			for r := 1; r <= n; r++ {
				if v := maths.Choose(n, r); v.GT(mill) {
					count++
				}
			}
		}
		o.Stdoutln(count)
	}, &ecmodels.Execution{
		Want: "4075",
	})
}
