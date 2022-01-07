package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P57() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=57"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			num, den := maths.NewInt(3), maths.NewInt(2)
			var count int
			for _ = range maths.Range(1000) {
				tmp := den
				den = den.Plus(num)
				num = tmp.Times(maths.NewInt(2)).Plus(num)
				if len(num.Digits()) > len(den.Digits()) {
					count++
				}
			}
			o.Stdoutln(count)
		}),
	)
}
