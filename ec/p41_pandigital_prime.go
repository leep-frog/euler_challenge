package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P41() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=41"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			best := 0
			max := 1_000_000_0
			p := generator.Primes()
			for pn := p.Next(); pn < max; pn = p.Next() {
				o.Stdoutln(pn)
				if maths.Pandigital(pn) {
					best = pn
					o.Stdoutln(best)
					fmt.Println(best)
				}
			}
			o.Stdoutln(best)
		}),
	)
}
