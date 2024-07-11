package p357

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P357() *ecmodels.Problem {
	return ecmodels.NoInputNode(357, func(o command.Output) {
		o.Stdoutln(primeGeneratingIntegers(0, 1, 100_000_000, generator.Primes()))
	}, &ecmodels.Execution{
		Want:     "1739023853137",
		Estimate: 45,
	})
}

func primeGeneratingIntegers(idx, val, max int, p *generator.Prime) *maths.Int {
	if val > max {
		return maths.Zero()
	}

	count := maths.Zero()
	if divisorSumsArePrime(val, p) {
		count = count.PlusInt(val)
	}

	// Numbers with multiple of the same prime factor cannot be included:
	// k = p_a ^ 2 * ...
	// Factors: k = (p_a * X) and (p_a * Y)
	// Sum:     p_a * X + p_a * Y = p_a (X + Y) which is divisible by p_a

	// Given the above, iterate over the next prime to use exactly once
	for i := idx; p.Nth(i)*val <= max && p.Nth(i) < max/2; i++ {
		count = count.Plus(primeGeneratingIntegers(i+1, p.Nth(i)*val, max, p))
	}
	return count
}

func divisorSumsArePrime(k int, p *generator.Prime) bool {

	if k%2 == 1 && k != 1 {
		return false
	}

	if !p.Contains(k + 1) {
		return false
	}

	for _, f := range p.Factors(k) {
		v := f + (k / f)
		if !p.Contains(v) {
			return false
		}
	}
	return true
}
