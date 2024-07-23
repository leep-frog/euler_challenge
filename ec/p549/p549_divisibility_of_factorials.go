package p549

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

// Related to p320
func P549() *ecmodels.Problem {
	return ecmodels.IntInputNode(549, func(o command.Output, n int) {
		p := generator.Primes()
		cache := map[int]int{}

		var sum int
		for i := 2; i <= maths.Pow(10, n); i++ {
			sum += solveIt(p, cache, i)
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "2012",
		},
		{
			Args:     []string{"8"},
			Want:     "476001479068717",
			Estimate: 75,
		},
	})
}

func solveIt(p *generator.Prime, cache map[int]int, n int) int {
	return generator.CompositeCacher[int](p, n, cache,
		func(i int) int {
			return i
		},
		func(primeFactor, otherFactor int) int {

			largest := cache[otherFactor]

			cnt := 1
			for ; otherFactor > 1 && otherFactor%primeFactor == 0; cnt, otherFactor = cnt+1, otherFactor/primeFactor {
			}
			res := factorCount(primeFactor, cnt)

			if res > largest {
				return res
			}
			return largest
		})
}

// func smallS(n int, p *generator.Prime) int {
// 	b := maths.Largest[int, int]()
// 	fs := p.PrimeFactors(n)
// 	for f, cnt := range fs {
// 		b.Check(factorCount(f, cnt))
// 	}
// 	return b.Best()
// }

func factorCount(f int, cnt int) int {
	// Assume f > 1 and cnt > 0

	// Get the largest power that is less than cnt
	power := f
	factorCount := 1
	for factorCount <= cnt {
		power *= f
		factorCount = factorCount*f + 1
	}

	// Set the number we are currently at
	var number int

	for cnt > 0 {
		// Decrement power
		power /= f
		factorCount = (factorCount - 1) / f

		times := cnt / factorCount
		number += power * times
		cnt = cnt % factorCount
	}

	return number
}
