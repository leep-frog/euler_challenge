package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P47() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=47"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			p := generator.Primes()
			var row int
			for i := 1; ; i++ {
				if len(generator.PrimeFactors(i, p)) >= n {
					row++
					if row == n {
						o.Stdoutln(i - (n - 1))
						return
					}
				} else {
					row = 0
				}
			}
		}),
	)
}
