package p401

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

// type to limit scope of functions
type problem401 struct{}

// Returns (1^2 + 2^2 + 3^2 + ... + k^2) % mod
func (*problem401) comb(k, mod int) int {
	a, b, c := k, k+1, 2*k+1
	if a%2 == 0 {
		a /= 2
	} else {
		b /= 2
	}
	if a%3 == 0 {
		a /= 3
	} else if b%3 == 0 {
		b /= 3
	} else {
		c /= 3
	}

	return (((a * b) % mod) * c) % mod
}

func (p *problem401) sigma2(n int) int {
	mod := 1_000_000_000
	var sum int
	for k := 1; k <= maths.Sqrt(n); k++ {
		t := n / k
		k, t = k%mod, t%mod

		// Divisor pairs included already: (1, k), (2, k), ...,
		// Divisor pairs not included yet: (k, k), (k+1, k), ... (t, k)
		// This breaks into:
		// S1: [k, k,   k, ...] length t-k
		// S2: [k, k+1, k+2, ..., t]

		// S1:
		sum += ((t - k) * ((k * k) % mod)) % mod

		// Other divisors created are 1, 2, 3, ..., t
		// Sum is 1, 4, 9, ..., t
		// = t(t+1)(2t+1)/6
		// However, we already accounted for all parts of t that are less than k
		// so we want to remove 1, 2, 3, ..., k
		// Let C(a, b) = a^2 + (a+1)^2 + (a+2)^2 + ... + b^2
		// S2 = C(k, t)
		//    = C(k, t) + C(1, k-1) - C(1, k-1)
		//    = [     C(1, t)     ] - C(1, k-1)
		sum = (sum + p.comb(t, mod) - p.comb(k-1, mod)) % mod
	}
	return sum
}

// Used to verify elegant solution for smaller values of n
func (*problem401) bruteForce(n int) int {
	p := generator.Primes()
	var sum int
	for k := 1; k <= n; k++ {
		for _, f := range p.Factors(k) {
			sum += f * f
		}
	}
	return sum
}

func P401() *ecmodels.Problem {
	return ecmodels.IntInputNode(401, func(o command.Output, n int) {
		p := &problem401{}
		o.Stdoutln(p.sigma2(n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"6"},
			Want: "113",
		},
		{
			Args: []string{"1_000_000_000_000_000"},
			Want: "281632621",
		},
	})
}
