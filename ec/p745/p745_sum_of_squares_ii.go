package p745

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

const mod = 1_000_000_007

func P745() *ecmodels.Problem {
	return ecmodels.IntInputNode(745, func(o command.Output, n int) {
		o.Stdoutln(clever(maths.Pow(10, n)))
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "24",
		},
		{
			Args: []string{"2"},
			Want: "767",
		},
		{
			Args:     []string{"14"},
			Want:     "94586478",
			Estimate: 150,
		},
	})
}

// Assume sum is the sum of G(n), but where it's when only numbers
// from 1 through k are considered. When considering (k+1), simply consider the numbers to remove
// For primes, we simply subtract 1 (for 3^2 ==> + 9 - 1)
// For composites, however, we need to consider overlaps:
// (2*3)^2 ==> Subtract 4 and 9 since those are no longer best for this spot
// But then also consider that both of those values subtracted a 1, so need to add one back
// So on and so forth with extra including/excluding at different spots.
func clever(n int) int {
	p := generator.Primes()

	sum := n

	mapper := make([]map[int]int, maths.Sqrt(n)+1, maths.Sqrt(n)+1)

	for k := 2; k*k <= n; k++ {

		// The number of times we see k
		t := n / (k * k)

		// Only subtract ones
		if p.Contains(k) {
			sum = (sum + t*(k*k-1)) % mod
			mapper[k] = map[int]int{
				// k * k: 1,
				1: -1,
			}
			continue
		}

		offset := k * k
		pfs := p.PrimeFactors(k)

		nm, fm := map[int]int{}, map[int]int{}

		for f := range pfs {
			c := k / f
			for k, v := range mapper[c] {
				nm[k] -= v
			}

			fm[c*c]--
			offset -= c * c
		}

		for sq, v := range nm {
			if v > 1 {
				offset += sq
				fm[sq] = 1
			} else if v < -1 {
				offset -= sq
				fm[sq] = -1
			}
		}

		for _, k := range fm {
			if fm[k] == 0 {
				delete(fm, k)
			}
		}

		mapper[k] = fm

		sum = (sum + t*(offset)) % mod
	}
	return sum
}

func brute(n int) int {
	var sum int
	p := generator.Primes()
	for i := 1; i <= n; i++ {
		coef := 1
		for f, cnt := range p.PrimeFactors(i) {
			coef = coef * maths.Pow(f, cnt-(cnt%2))
		}
		sum += coef
	}
	return sum
}
