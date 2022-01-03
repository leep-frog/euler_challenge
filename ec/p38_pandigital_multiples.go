package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P38() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=38"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			best := maths.Largest()
			for i := 1; i < 1000000; i++ {
				m := map[int]bool{}

				var vs int
				for n := 1; ; n++ {
					for _, d := range maths.Digits(n * i) {
						if m[d] || d == 0 {
							goto NEXT
						}
						m[d] = true
						vs = vs*10 + (d % 10)
					}

					// Check if addition
					if len(m) == 9 {
						best.IndexCheck(i, vs)
						goto NEXT
					}
				}
			NEXT:
			}
			o.Stdoutln(best.Best(), best.BestIndex())
		}),
	)
}
