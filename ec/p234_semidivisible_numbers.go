package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P234() *problem {
	return intInputNode(234, func(o command.Output, n int) {
		prev := 2
		sum := maths.NewInt(0)
		for g, p := generator.Primes().Start(1); prev*prev <= n; p = g.Next() {
			a1 := maths.NewInt(int64(semidivSums(n, prev*prev, p*p, prev)))
			a2 := maths.NewInt(int64(semidivSums(n, prev*prev, p*p, p)))
			a3 := maths.NewInt(int64(2 * prev * p))
			if a3.GT(maths.NewInt(int64(2 * n))) {
				a3 = maths.NewInt(0)
			}
			sum = sum.Plus(a1).Plus(a2).Minus(a3)
			prev = p
		}
		o.Stdoutln(sum)
	})
}

// Return the sum of numbers between start and end (exclusive)
// that are divisble by den
func semidivSums(n, start, end, den int) int {
	s := (start / den) + 1
	e := maths.Min(n, end-1) / den

	if s > e {
		return 0
	}

	// Get sum of startIdx ... endIdx (s ... e)
	// s + ... + e = s * (e - s + 1) + 1 + 2 + ... (e - s)
	//             = s * (e - s + 1) + (e - s)*(e - s + 1)/2
	// Then multiply by den
	return den * (s*(e-s+1) + (e-s)*(e-s+1)/2)
}
