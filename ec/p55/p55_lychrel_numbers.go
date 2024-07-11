package p55

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P55() *ecmodels.Problem {
	return ecmodels.NoInputNode(55, func(o command.Output) {
		var count int
		for i := 1; i < 10_000; i++ {
			big := maths.NewInt(i)
			big = big.Plus(big.Reverse())
			for j := 0; j < 49; j++ {
				if big.Palindrome() {
					goto NOPE
				}
				big = big.Plus(big.Reverse())
			}
			count++
		NOPE:
		}
		o.Stdoutln(count)
	}, &ecmodels.Execution{
		Want: "249",
	})
}
