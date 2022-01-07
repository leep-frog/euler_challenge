package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P56() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=56"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			best := maths.Largest()
			for a := 1; a < 100; a++ {
				for b := 1; b < 100; b++ {
					best.Check(maths.BigPow(a, b).DigitSum())
				}
			}
			o.Stdoutln(best.Best())
		}),
	)
}
