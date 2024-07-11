package p234

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P234() *ecmodels.Problem {
	return ecmodels.IntInputNode(234, func(o command.Output, n int) {
		prev := 2
		sum := maths.NewInt(0)
		for g, p := generator.Primes().Start(1); prev*prev <= n; p = g.Next() {
			// prev^2 < p*prev < p^2 < 2*p*prev
			// Therefore the only semidivisble numbers are all of the numbers
			// between prev^2 and p^2 that are divisble by prev OR p EXCEPT prev*p.
			// Since [2*prev*p > p*p], we know that prev*p is the ONLY number that
			// is divisble by both prev and p, and is between prev^2 and p^2.

			// All of the numbers divisble by prev between prev^2 and p^2 (including prev*p)
			a1 := maths.NewInt(semidivSums(n, prev*prev, p*p, prev))
			// All of the numbers divisble by p    between prev^2 and p^2 (including prev*p)
			a2 := maths.NewInt(semidivSums(n, prev*prev, p*p, p))
			// Subtract prev*p twice ONLY if it is less than n (if it is greater than
			// n, then semidivSums will not have considered it in its return value).
			a3 := maths.NewInt(2 * prev * p)
			if a3.GT(maths.NewInt(2 * n)) {
				a3 = maths.NewInt(0)
			}

			// Add up the semidivisble values
			sum = sum.Plus(a1).Plus(a2).Minus(a3)

			prev = p
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"999966663332"},
			Want: "1259187438574927161",
		},
		{
			Args: []string{"1000"},
			Want: "34825",
		},
		{
			Args: []string{"15"},
			Want: "30",
		},
	})
}

// Return the sum of numbers between start and end (exclusive)
// that are divisble by 'den'.
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
