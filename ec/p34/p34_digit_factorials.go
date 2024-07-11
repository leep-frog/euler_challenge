package p34

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P34() *ecmodels.Problem {
	return ecmodels.NoInputNode(34, func(o command.Output) {
		fs := map[int]int{
			0: 1,
		}
		for i := 1; i <= 9; i++ {
			fs[i] = maths.FactorialI(i)
		}

		var superSum int
		for i := 3; i < 100_000; i++ {
			var sum int
			for c := i; c > 0; c /= 10 {
				sum += fs[c%10]
			}
			if sum == i {
				superSum += i
			}
		}
		o.Stdoutln(superSum)
	}, &ecmodels.Execution{
		Want: "40730",
	})
}
