package p45

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P45() *ecmodels.Problem {
	return ecmodels.NoInputNode(45, func(o command.Output) {
		for iter, hn := generator.Hexagonals().Start(0); ; hn = iter.Next() {
			if hn <= 40755 {
				continue
			}
			if generator.IsPentagonal(hn) && generator.IsTriangular(hn) {
				o.Stdoutln(hn)
				return
			}
		}
	}, &ecmodels.Execution{
		Want: "1533776805",
	})
}
