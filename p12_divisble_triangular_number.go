package main

import (
	"github.com/leep-frog/command"
)

func p12() *command.Node {
	return command.SerialNodes(
		command.Description("Find a triangular number with at least N divisors"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			primer := &Primer{}
			for i := 0; i < d.Int(N); i++ {
				primer.Next()
			}
			o.Stdoutf("%d", primer.Last())
			return nil
		}),
	)
}
