package p808

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P808() *ecmodels.Problem {
	return ecmodels.IntInputNode(808, func(o command.Output, n int) {
		primes := generator.Primes()
		g := primes.Iterator()
		var count, sum int
		for p := g.Next(); count < n; p = g.Next() {
			square := p * p
			if maths.Palindrome(square) {
				continue
			}
			reverse := maths.Reverse(square)
			if maths.IsSquare(reverse) && primes.Contains(maths.IntSquareRoot(reverse)) {
				sum += square
				count++
			}
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"50"},
			Want:     "3807504276997394",
			Estimate: 10,
		},
	})
}
