package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P39() *problem {
	return noInputNode(39, func(o command.Output) {
		best := maths.LargestIncremental()
		for p := 1; p <= 1000; p++ {
			for a := 1; a < p; a++ {
				for b := a + 1; a+b < p; b++ {
					if c2 := a*a + b*b; maths.IsSquare(c2) {
						best.Increment(a + b + maths.Sqrt(c2))
					}
				}
			}
		}
		o.Stdoutln(best.BestIndex())
	}, &execution{
		want:     "840",
		estimate: 0.25,
	})
}
