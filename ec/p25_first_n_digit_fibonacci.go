package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P25() *command.Node {
	return command.SerialNodes(
		command.Description("Find the first fibonacci digit with n digits"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			for g, i := generator.BigFibonaccis(), 1; ; i++ {
				o.Stdoutln(i)
				if len(g.Nth(i).String()) >= n {
					o.Stdoutln(i + 1)
					return
				}
			}
		}),
	)
}
