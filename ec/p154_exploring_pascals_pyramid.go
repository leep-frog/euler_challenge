package eulerchallenge

import (
	"github.com/leep-frog/command"
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
*/

func P154() *problem {
	return intsInputNode(154, 2, 0, func(o command.Output, ns []int) {
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
			coef := fc.Choose(n, i)
			for j := offset; j <= i/2; j++ {
				// See if it is divisble
				divisble := true
				v := fc.Choose(i, j)
				for i, mn := range minNeeded {
					if coef[i]+v[i] < mn {
						divisble = false
					}
				}
				if !divisble {
					continue
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
			}
		}

		o.Stdoutln(cnt)
	}, []*execution{
		{
			args: []string{"4", "4"},
			want: "9",
		},
		{
			args: []string{"4", "2"},
			want: "12",
		},
		{
			args: []string{"5", "5"},
			want: "18",
		},
		{
			args: []string{"5", "2"},
			want: "12",
		},
		{
			args: []string{"5", "4"},
			want: "3",
		},
		{
			args: []string{"5", "10"},
			want: "12",
		},
		{
			args:     []string{"200000", "1000000000000"},
			want:     "479742450",
			estimate: 300,
		},
	})
}

type FactorialChecker struct {
	factors []int

	divs [][]int
}

// Returns the number of each factor in the numerator
func (fc *FactorialChecker) Choose(n, k int) []int {
	// n! / (k!(n-k)!)

	var rs []int
	for i := range fc.factors {
		// Number of the factors in n!
		nf := fc.divs[n][i]
		// Number of the factors in k!
		kf := fc.divs[k][i]
		// Number of the factors in k!
		knf := fc.divs[n-k][i]

		// Number of factors in the numerator:
		rs = append(rs, nf-kf-knf)
	}
	return rs
}

// Note: only works if factors are primes.
func NewFC(k int, factors []int) *FactorialChecker {
	// currentF is an array of [factor, factor^x] where factor^x is
	// bigger than the current number we are on in the iteration below.
	var currentF [][]int
	for _, f := range factors {
		currentF = append(currentF, []int{f, f})
	}

	lf := len(factors)
	divs := [][]int{make([]int, lf, lf)}
	for i := 1; i <= k; i++ {
		var newRow []int
		for ci, cf := range currentF {
			v := divs[i-1][ci]
			for s := i; s%cf[1] == 0; s, v = s/cf[1], v+1 {
			}
			newRow = append(newRow, v)
		}
		divs = append(divs, newRow)
	}
	return &FactorialChecker{factors, divs}
}
