package p757

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/slices"
)

func P757() *ecmodels.Problem {
	return ecmodels.IntInputNode(757, func(o command.Output, n int) {
		o.Stdoutln(mediumCleverr(maths.Pow(10, n)))
	}, []*ecmodels.Execution{
		{
			Args: []string{"6"},
			Want: "2851",
		},
		{
			Args:     []string{"14"},
			Want:     "75737353",
			Estimate: 7.5,
		},
	})
}

// Realized that all answers are of the form a*(a+1)*b*(b+1)
// For example, 84 = 1*2*6*7
func mediumCleverr(n int) int {
	// By using a slice instead of a set (map[int]bool), we go from 16s to 6s
	var m []int
	for a := 1; a*(a+1)*a*(a+1) <= n; a++ {
		for b := a; a*(a+1)*b*(b+1) <= n; b++ {
			m = append(m, a*(a+1)*b*(b+1))
		}
	}

	// Count the unique elements in the slice
	slices.Sort(m)
	prev := -1
	var cnt int
	for _, v := range m {
		if v != prev {
			cnt++
			prev = v
		}
	}
	return cnt
}

// This doesn't work because some numbers overlap
// e.g. 144 = 1*2*8*9 = 3*4*3*4
// Could probably negate the overlap pattern, but the medium-clever solution
// works quickly enough
func clever(n int) int {
	var sum int
	for num := 1; num*(num+1)*num*(num+1) <= n; num++ {
		a, b := num, num+1
		// Find the number x, such that a*b*x*(x+1) < n
		//                              x*(x+1) < n/(a*b)
		xLower := maths.Sqrt(n/(a*b)) - 1

		for (xLower+1)*(xLower+2)*a*b <= n {
			xLower++
		}
		sum += (xLower - num + 1)
	}
	return sum
}

func brute(n int) int {
	p := generator.Primes()
	var cnt int
	for k := 2; k <= n; k++ {
		var stealthy bool

		m := map[int]bool{}
		for _, a := range p.Factors(k) {
			b := k / a
			if a > b {
				continue
			}

			m[a+b] = true
			stealthy = stealthy || m[a+b+1] || m[a+b-1]
		}
		if stealthy {
			cnt++
		}
	}
	return cnt
}
