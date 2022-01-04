package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P44() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=44"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			best := -1
			pents := generator.Pentagonals()
			for i := 0; ; i++ {
				pi := pents.Nth(i)
				for j := i - 1; j >= 0 && (best == -1 || pi-pents.Nth(j) < best); j-- {
					pj := pents.Nth(j)
					if generator.IsPentagonal(pi-pj) && generator.IsPentagonal(pi+pj) && (best == -1 || pi-pj < best) {
						best = pi - pj
					}
				}
				// We can't do better if the next difference is bigger than the current best.
				if best != -1 && pents.Nth(i+1)-pents.Nth(i) >= best {
					break
				}
			}
			o.Stdoutln(best)
		}),
	)
}
