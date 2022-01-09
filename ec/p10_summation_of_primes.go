package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P10() *command.Node {
	return command.SerialNodes(
		command.Description("Find the sum of all primes lower than N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			p := generator.Primes()

			var sum int
			for pn := p.Next(); pn < d.Int(N); pn = p.Next() {
				sum += pn
			}
			o.Stdoutln(sum)
		}),
	)
}
