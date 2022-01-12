package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P64() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=64"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)
			var count int
			for k := 2; k <= n; k++ {
				_, period := maths.SquareRootPeriod(k)
				if len(period)%2 == 1 {
					count++
				}
			}
			o.Stdoutln(count)
		}),
	)
}
