package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P63() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=63"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			var count int
			for pow := 1; pow < 50; pow++ {
				n := 1
				for ; len(maths.BigPow(n, pow).String()) < pow; n++ {
				}
				for ; len(maths.BigPow(n, pow).String()) == pow; n++ {
					count++
				}
			}
			o.Stdoutln(count)
		}),
	)
}
