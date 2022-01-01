package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P27() *command.Node {
	return command.SerialNodes(
		command.Description("Find the quadratic that produces the largest set of consecutive primes"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			p := generator.Primes()

			var max, maxI int
			for a := -n + 1; a < n; a++ {
				for b := -n; b <= n; b++ {
					// Try positive direction
					k := 0
					for ; p.Contains(k*k + a*k + b); k++ {
					}
					if k > max {
						max = k
						maxI = a * b
					}

					// Try negative direction
					k = 0
					for ; p.Contains(k*k + a*k + b); k-- {
					}
					if k > max {
						max = k
						maxI = a * b
					}
				}
			}
			o.Stdoutln(maxI)
		}),
	)
}
