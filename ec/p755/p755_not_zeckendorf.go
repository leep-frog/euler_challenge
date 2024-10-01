package p755

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

var (
	f = generator.Fibonaccis()
)

func P755() *ecmodels.Problem {
	return ecmodels.IntInputNode(755, func(o command.Output, n int) {
		o.Stdoutln(fibsLessThan(n, -1))
	}, []*ecmodels.Execution{
		{
			Args: []string{"100"},
			Want: "415",
		},
		{
			Args: []string{"10_000"},
			Want: "312807",
		},
		{
			Args: []string{"10_000_000_000_000"},
			Want: "2877071595975576960",
		},
	})
}

// fibsLessThan calulates f(n) by iterating over each Fibonacci number, x, and
// considering which numbers less than n can be made with x as the largest Fibonacci number used.
func fibsLessThan(n, upToFib int) int {

	cnt := 1 // f(0) = 1

	// First, consider all Fibonacci numbers where all permutations of the smaller numbers can be used
	var fibSum int
	fibIdx := 1
	for ; (fibIdx <= upToFib || upToFib < 0) && fibSum+f.Nth(fibIdx) <= n; fibIdx++ {
		fibSum += f.Nth(fibIdx)
		// Include all numbers you can create that include this fibonacci number as the largest fibonacci number
		cnt += maths.Pow(2, fibIdx-1)
	}

	// For numbers where the fibSum > n, we need a bit more nuance
	for ; (fibIdx <= upToFib || upToFib < 0) && f.Nth(fibIdx) <= n; fibIdx++ {
		// Calculate the numbers less than the remainder using only smaller Fibonaccis
		cnt += fibsLessThan(n-f.Nth(fibIdx), fibIdx-1)
	}
	return cnt
}
