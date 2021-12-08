package main

import (
	"github.com/leep-frog/command"
)

func p7() *command.Node {
	return command.SerialNodes(
		command.Description("Find the Nth prime number"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			primer := &Primer{}
			o.Stdoutf("%d", primer.Nth(d.Int(N)-1))
			return nil
		}),
	)
}
