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

		var sum int
		for i := 2; i <= maths.Pow(10, n); i++ {
			sum += smallS(i, p)
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "2012",
		},
		{
			Args: []string{"8"},
			Want: "476001479068717",
		},
	})
}

func smallS(n int, p *generator.Prime) int {
	b := maths.Largest[int, int]()
	fs := p.PrimeFactors(n)
	for f, cnt := range fs {
		b.Check(factorCount(f, cnt))
	}
	return b.Best()
}

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
