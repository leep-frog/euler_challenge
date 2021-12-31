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

			for i := 0; p.Next() < d.Int(N); i++ {
			}

			sum := 0
			for i := 0; i < p.Len()-1; i++ {
				sum += p.Nth(i)
			}
			o.Stdoutln(sum)
		}),
	)
}
