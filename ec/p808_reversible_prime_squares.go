package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P808() *problem {
	return intInputNode(808, func(o command.Output, n int) {
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
	}, []*execution{
		{
			args:     []string{"50"},
			want:     "3807504276997394",
			estimate: 10,
		},
	})
}
