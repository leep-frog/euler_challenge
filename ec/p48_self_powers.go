package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P48() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=48"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			res := 0
			largest := 10_000_000_000
			_ = largest
			for i := 1; i <= n; i++ {
				prod := i
				for j := 1; j < i; j++ {
					prod = (prod * i) % largest
				}
				res = (res + prod) % largest
			}
			o.Stdoutln(res % largest)
		}),
	)
}
