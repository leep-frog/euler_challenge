package eulerchallenge

import (
	"fmt"
	"time"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

type problem501 struct{}

// a^3 * b^1
func (p *problem501) SingleAndCubePrimeCount(n int, primes *generator.Prime) int {
	var sum int
	for iter, prime := primes.Start(0); 2*maths.Pow(prime, 3) <= n; prime = iter.Next() {
		fmt.Println("1a", prime, time.Now())
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

func P501() *problem {
	return intInputNode(501, func(o command.Output, n int) {
		primes := generator.BatchedSievedPrimes()
		p := &problem501{}
		o.Stdoutln(p.SingleAndCubePrimeCount(n, primes), +p.ThreeDistinctPrimeCount(n, primes)+p.SeventhExpCount(n, primes))
	}, []*execution{
		{
			args: []string{"100"},
			want: "10",
		},
		{
			args: []string{"1000"},
			want: "180",
		},
		{
			args: []string{"1000000"},
			want: "224427",
		},
		// Takes over an hour. Other solvers used a built-in Mathematica function
		// that we don't have access to and implementation (https://en.wikipedia.org/wiki/Prime-counting_function)
		// is convoluted. I'm fine not implementing my own efficient PrimePi()
		// function and assuming we just would have had one.
		{
			args:     []string{"1_000_000_000_000"},
			want:     "197912312715",
			estimate: 75 * 60,
		},
	})
}
