package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P7() *command.Node {
	return command.SerialNodes(
		command.Description("Find the Nth prime number"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			o.Stdoutln(generator.Primes().Nth(d.Int(N) - 1))
		}),
	)
}
