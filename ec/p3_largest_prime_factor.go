package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

// TODO: move all of these to helper directory

func P3() *command.Node {
	return command.SerialNodes(
		command.Description("Find the largest prime factor of N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			factors := generator.PrimeFactors(d.Int(N), generator.Primes())

			max := 0
			for f := range factors {
				if f > max {
					max = f
				}
			}

			o.Stdoutln(max)
		}),
	)
}
