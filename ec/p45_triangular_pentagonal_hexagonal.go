package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P45() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=45"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			h := generator.Hexagonals()
			for hn := h.Next(); ; hn = h.Next() {
				if hn <= 40755 {
					continue
				}
				if generator.IsPentagonal(hn) && generator.IsTriangular(hn) {
					o.Stdoutln(hn)
					return
				}
			}
		}),
	)
}
