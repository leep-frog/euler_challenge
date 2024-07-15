package p500

import (
	"math"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P500() *ecmodels.Problem {
	return ecmodels.IntInputNode(500, func(o command.Output, n int) {

		p := generator.Primes()

		// The pattern (for powers of 2 bests) adds the smallest prime or addition of p^(2^n-1) to the previous solution:
		// Number, num factors, prime factors
		//     2     2           map[2:1]                   Add 2
		//     6     4           map[2:1 3:1]               Add 3
		//    24     8           map[2:3 3:1]               Add 2^2 = 4
		//   120    16           map[2:3 3:1 5:1]           Add 5
		//   840    32           map[2:3 3:1 5:1 7:1]       Add 7
		//  7560    64           map[2:3 3:3 5:1 7:1]       Add 3^3 = 9
		// 83160   128           map[2:3 3:3 5:1 7:1 11:1]  Add 11

		// Make a heap wher the smallest element is the next of the above sequence
		heap := maths.NewHeap(func(a, b *heapElement) bool {
			return float64(a.exp)*math.Log(float64(a.v)) < float64(b.exp)*math.Log(float64(b.v))
		})

		// Add the initial element
		heap.Push(&heapElement{
			v:            2,
			exp:          1,
			nextPrimeIdx: 1,
		})

		fs := map[int]int{}
		for i := 0; i < n; i++ {
			el := heap.Pop()
			fs[el.v] += el.exp

			if el.nextPrimeIdx > 0 {
				heap.Push(&heapElement{
					v:            p.Nth(el.nextPrimeIdx),
					exp:          1,
					nextPrimeIdx: el.nextPrimeIdx + 1,
				})
			}

			heap.Push(&heapElement{
				v:   el.v,
				exp: el.exp * 2,
			})
		}

		v := maths.Zero()
		for _, c := range fs {
			v = v.PlusInt(c)
		}

		r := maths.One()
		for k, v := range fs {
			coef := maths.Pow(k, v)
			r = r.TimesInt(coef)
			r = r.Mod(maths.NewInt(500500507))
		}

		o.Stdoutln(r)
	}, []*ecmodels.Execution{
		{
			Args: []string{"4"},
			Want: "120",
		},
		{
			Args:     []string{"500500"},
			Want:     "35407281",
			Estimate: 2,
		},
	})
}

type heapElement struct {
	v            int
	exp          int
	nextPrimeIdx int
}
