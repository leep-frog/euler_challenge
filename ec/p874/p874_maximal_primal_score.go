package p874

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

var (
	p       = generator.Primes()
	cache   = map[string]int{}
	absBest = maths.Largest[int, int]()
)

func P874() *ecmodels.Problem {
	return ecmodels.IntInputNode(874, func(o command.Output, n int) {

		prime := p.Nth(n) // aka count

		max := prime * (n - 1)

		negativeSpace := max % n

		o.Stdoutln(dp(negativeSpace, negativeSpace, prime, n, 0))

	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "14",
		},
		{
			Args: []string{"4"},
			Want: "75",
		},
		{
			Args: []string{"7000"},
			Want: "4992775389",
			Skip: "Found value in incremental execution, but full execution takes much longer (never completed so not sure how long tbh)",
		},
	})
}

// The solution does a two-fold approach:
//  1. dynamic programming to speed things up (this builds the sum from bottom up).
//     However, this only checks values when we reach another root node which can take a while.
//  2. building the currentScore sum from top down to print out the current best at each
//     leaf (to get around the problem from part 1.
//
// Based on the thread, we can also do the knapsack solution or check for optimal decrements (-4, and -9),
// but the way I solved it is fine enough.
//
// If we find more uses for a generic knapsack solver, I can implement that,
// otherwise, leave as is.
func dp(negativeSpace, at, numbersRemaining, n, currentScore int) int {
	code := fmt.Sprintf("%d-%d-%d", negativeSpace, at, numbersRemaining)
	if v, ok := cache[code]; ok {
		if n == 7000 && absBest.Check(currentScore+v) {
			fmt.Println("NEW BESTER", currentScore+v)
		}
		return v
	}

	if negativeSpace == 0 {
		r := numbersRemaining * p.Nth(n-1)
		cache[code] = r
		if n == 7000 && absBest.Check(currentScore+r) {
			fmt.Println("NEW BEST", currentScore+r)
		}
		return r
	}

	best := maths.Largest[int, int]()

	// Check without adding another of these
	if at > 1 {
		best.Check(dp(negativeSpace, at-1, numbersRemaining, n, currentScore))
	}

	// Check with adding one of these
	best.Check(p.Nth(n-1-at) + dp(negativeSpace-at, maths.Min(at, negativeSpace-at), numbersRemaining-1, n, currentScore+p.Nth(n-1-at)))

	cache[code] = best.Best()
	return best.Best()
}
