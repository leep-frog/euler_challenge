package p501

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

type problem501 struct{}

// a^3 * b^1
func (p *problem501) SingleAndCubePrimeCount(n int, primes *generator.Prime) int {
	var sum int
	for iter, prime := primes.Start(0); 2*maths.Pow(prime, 3) <= n; prime = iter.Next() {
		sum += primes.PrimePi(n / maths.Pow(prime, 3))
		// Can also check iter
		if maths.Pow(prime, 4) <= n {
			sum--
		}
	}
	return sum
}

// a*b*c
func (p *problem501) ThreeDistinctPrimeCount(n int, primes *generator.Prime) int {
	var sum int

	// Do every pair of primes
	for pi := 0; primes.Nth(pi)*primes.Nth(pi+1)*primes.Nth(pi+2) <= n; pi++ {
		pRem := n / primes.Nth(pi)
		for qi := pi + 1; primes.Nth(qi)*primes.Nth(qi+1) <= pRem; qi++ {
			pqRem := pRem / primes.Nth(qi)
			primeCnt := primes.PrimePi(pqRem)
			// Only count numbers greater than qi
			sum += maths.Max(0, primeCnt-qi-1)
		}
	}

	return sum
}

// a^7
func (p *problem501) SeventhExpCount(n int, primes *generator.Prime) int {
	var sum int
	for g, p := primes.Start(0); maths.Pow(p, 7) <= n; p = g.Next() {
		sum++
	}
	return sum
}

func P501() *ecmodels.Problem {
	return ecmodels.IntInputNode(501, func(o command.Output, n int) {
		primes := generator.BatchedSievedPrimes()
		p := &problem501{}
		o.Stdoutln(p.SingleAndCubePrimeCount(n, primes) + p.ThreeDistinctPrimeCount(n, primes) + p.SeventhExpCount(n, primes))
	}, []*ecmodels.Execution{
		{
			Args:     []string{"100"},
			Want:     "10",
			Estimate: 0.5,
		},
		{
			Args:     []string{"1000"},
			Want:     "180",
			Estimate: 0.5,
		},
		{
			Args:     []string{"1000000"},
			Want:     "224427",
			Estimate: 2.5,
		},
		{
			Args: []string{"1_000_000_000_000"},
			Want: "197912312715",
			Skip: strings.Join([]string{
				`Takes over an hour. Other solvers used a built-in Mathematica function`,
				`that we don't have access to and implementation (https://en.wikipedia.org/wiki/Prime-counting_function)`,
				`is convoluted. I'm fine not implementing my own efficient PrimePi()`,
				`function and assuming we just would have had one.`,
			}, "\n"),
			Estimate: 75 * 60,
		},
	})
}
