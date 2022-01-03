package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P41() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=41"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			best := 0
			/*for p := generator.Primes(); p.Next() < 100_000; {
				if maths.Pandigital(p.Last()) {
					//best = p.Last()
					//o.Stdoutln(best)
					fmt.Println(best)
				}
			}*/
			o.Stdoutln(best)
		}),
	)
}
