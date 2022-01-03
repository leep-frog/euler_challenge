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
			for p := generator.Primes(); p.Next() < 1_000; {
				if maths.Pandigital(p.Last()) {
					best = p.Last()
					//o.Stdoutln(best)
					fmt.Println(best)
				}
			}
			o.Stdoutln(best)
		}),
	)
}
