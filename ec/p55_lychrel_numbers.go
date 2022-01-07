package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P55() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=55"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			var count int
			for i := 1; i < 10_000; i++ {
				big := maths.NewInt(int64(i))
				big = big.Plus(big.Reverse())
				for j := 0; j < 49; j++ {
					if big.Palindrome() {
						goto NOPE
						break
					}
					big = big.Plus(big.Reverse())
				}
				count++
			NOPE:
			}
			o.Stdoutln(count)
		}),
	)
}
