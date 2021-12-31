package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P2() *command.Node {
	return command.SerialNodes(
		command.Description("Find the sum of all even fibonacci numbers less than N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			fibs := generator.Fibonaccis()
			var sum int
			for i := fibs.Next(); i < d.Int(N); i = fibs.Next() {
				if i%2 == 0 {
					sum += i
				}
			}
			o.Stdoutln(sum)
		}),
	)
}
