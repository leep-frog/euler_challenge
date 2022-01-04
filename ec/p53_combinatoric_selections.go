package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P53() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=53"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
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
		}),
	)
}
