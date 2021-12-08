package main

import (
	"github.com/leep-frog/command"
)

func p2() *command.Node {
	return command.SerialNodes(
		command.Description("Find the sum of all even fibonacci numbers less than N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			fibs := Fibonaccis()
			var sum int
			for i := fibs.Next(); i < d.Int(N); i = fibs.Next() {
				if i%2 == 0 {
					sum += i
				}
			}
			o.Stdoutf("%d", sum)
			return nil
		}),
	)
}
