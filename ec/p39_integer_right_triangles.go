package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P39() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=39"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
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
		}),
	)
}
