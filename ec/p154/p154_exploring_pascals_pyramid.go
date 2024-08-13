package p154

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

/*
Our solution involves splitting the triangle into 6ths by drawing
the bisecting angle from all vertices. Any point in the open space,
by symmetry, exists 6 times in the triangle. Any point on the edge of
the 1/6th triangle exists 3 times, and any number in the middle exists
just once.

Example triangles

P_0:
1

P_1:
 1
1 1

P_2:
  1
 2 2
1 2 1

P_3:
   1
	3 3
 3 6 3
1 3 3 1

P_4:
      1
    4   4
   6  12  6
 4  12  12  4
1  4   6   4  1

P_5:
       1
      5 5
    10 20 10
  10 30 30 10
 5  20 30 20 5
1 5  10  10 5 1

I looked into optimization, but everyone else's solution also took on the
order of 1+ minutes so not much improvement to be made here.
*/

func P154() *ecmodels.Problem {
	return ecmodels.IntsInputNode(154, 2, 0, func(o command.Output, ns []int) {
		// Pascal Pyramid layer
		n := ns[0]
		// Divisble by
		div := ns[1]

		// Create the factors slice (each element is [factor, number of factors needed]).
		divFactors := generator.Primes().PrimeFactors(div)
		var minNeeded, factors []int
		for f, m := range divFactors {
			factors = append(factors, f)
			minNeeded = append(minNeeded, m)
		}

		fc := NewFC(n, factors)

		cnt := 0
		for i, offset := n, 0; offset <= n/2; i, offset = i-1, offset+1 {

			// We need (n choose i) * (i choose j)
			// n choose i -> n! / (i! * (n-i)!)
			// i choose j -> i! / (j! * (i-j)!)
			// (n choose i) * (i choose j) =  (n! / (i! * (n-i)!)) * (i! / (j! * (i-j)!))
			//                             = n! / ( (n-i)! * j! * (i-j)! )

			// currentFactors tracks the current amount of each factor
			var currentFactors []int
			for idx := range factors {
				// n! / (n-i)!
				currentFactors = append(currentFactors, fc.divs[n][idx]-fc.divs[n-i][idx])
			}

			for j := offset; j <= i/2; j++ {
				// See if it is divisble
				for idx, f := range fc.divs[j] {
					// Divide by (j!) and (i-j)!
					// and see if there are still enough factors left over to divide minNeeded
					if currentFactors[idx]-f-fc.divs[i-j][idx] < minNeeded[idx] {
						goto END_LOOP
					}
				}

				// If on the left line and the downward line, we are in the middle (only 1).
				if j == offset && i%2 == 0 && j == i/2 {
					cnt++
				} else if j == offset || (i%2 == 0 && j == i/2) {
					// If we are on the left line OR the downward line, then symmetry produces only 3.
					cnt += 3
				} else {
					// Otherwise, symmetry produces 6 (6 identical triangles created when
					// bisect the triangle from all three vertices).
					cnt += 6
				}

			END_LOOP:
			}
		}

		o.Stdoutln(cnt)
	}, []*ecmodels.Execution{
		{
			Args: []string{"4", "4"},
			Want: "9",
		},
		{
			Args: []string{"4", "2"},
			Want: "12",
		},
		{
			Args: []string{"5", "5"},
			Want: "18",
		},
		{
			Args: []string{"5", "2"},
			Want: "12",
		},
		{
			Args: []string{"5", "4"},
			Want: "3",
		},
		{
			Args: []string{"5", "10"},
			Want: "12",
		},
		{
			Args:     []string{"200000", "1000000000000"},
			Want:     "479742450",
			Estimate: 15,
		},
	})
}

type FactorialChecker struct {
	factors []int

	divs [][]int
}

// Note: only works if factors are primes.
func NewFC(k int, factors []int) *FactorialChecker {
	lf := len(factors)
	divs := [][]int{make([]int, lf)}
	for i := 1; i <= k; i++ {
		var newRow []int
		for ci, cf := range factors {
			v := divs[i-1][ci]
			for s := i; s%cf == 0; s, v = s/cf, v+1 {
			}
			newRow = append(newRow, v)
		}
		divs = append(divs, newRow)
	}
	return &FactorialChecker{factors, divs}
}
