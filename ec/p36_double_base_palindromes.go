package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P36() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=36"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			var palins []int
			for j, prod := 1, 1; prod < n; j++ {
				palins = append(palins, maths.Palindromes(j)...)
				prod *= 10
			}

			var sum int
			for _, palin := range palins {
				if maths.ToBinary(palin).Palindrome() {
					sum += palin
				}
			}
			o.Stdoutln(sum)
		}),
	)
}
