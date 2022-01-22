package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P45() *problem {
	return noInputNode(45, func(o command.Output) {
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
	})
}
