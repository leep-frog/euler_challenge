package main

import (
	"github.com/leep-frog/command"
)

func p2() *command.Node {
	return command.SerialNodes(
		command.Description("Find the sum of all even fibonacci numbers less than N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			var sum int
			for a, b := 1, 2; b < d.Int(N); {
				if b%2 == 0 {
					sum += b
				}
				tmp := a + b
				a = b
				b = tmp
			}
			o.Stdoutf("%d", sum)
			return nil
		}),
	)
}
