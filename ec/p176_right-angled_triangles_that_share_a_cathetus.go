package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

// Noticed the following pattern:
// f(x) -> f(2x) -> f(4x) makes a linear progression.
// Let that progression be defined:
// a = f(x) (where x % 2 = 0)
// d = f(2x) - f(x) = f(4x) - f(x)
// If we want to find a number k for which f(k) = n, then we can just iterate
// through linear progressions until we find one that contains n.
//
// Let all values be defined as (3^p1) * (5^p2) * (7^p3) * (11^p4) * ...
// where p1 >= p2 >= p3 >= ...
// Let a_k, d_k be the a and d values for some k
// If we add another prime with power q, then a_k and d_k are updated as follows:
// a_r = a_k * (2q + 1) + 1
// d_r = a_d * (2q + 1)
// So, we can just iterate over sets of prime powers (p1, p2, p3, ...),
// keeping track of the linear progression the sequence makes, and see
// which ones hit n (keeping track of the best value k that does so).
func recur176(n, rem, max, a, d int, product *maths.Int, cur []int, best *maths.Bester[[]int, *maths.Int], primes *generator.Prime) {
	if a > n {
		return
	}

	if best.Set() && product.GT(best.Best()) {
		return
	}

	// If the linear progress hits n, check the value that does it.
	if (n-a)%d == 0 {
		twoExp := (n - a) / d
		res := product.Times(maths.BigPow(2, twoExp+1))
		best.IndexCheck(bread.Copy(cur), res)
	}

	if rem == 0 {
		return
	}

	// Add all possible powers equal to or lower than the current lowest exponent in the sequence (keep decreasing size of exponents).
	for exp := max; exp > 0; exp-- {
		newProduct := product.Times(maths.BigPow(primes.Nth(len(cur)+1), exp))
		recur176(n, rem-1, exp, a*(2*exp+1)+exp, d*(2*exp+1), newProduct, append(cur, exp), best, primes)
	}
}

func P176() *problem {
	return intsInputNode(176, 1, command.UnboundedList, func(o command.Output, n []int) {
		rbest := maths.SmallestT[[]int, *maths.Int]()
		recur176(n[0], 30, 30, 0, 1, maths.One(), nil, rbest, generator.Primes())
		o.Stdoutln(rbest.Best())
	}, []*execution{
		{
			args: []string{"4"},
			want: "12",
		},
		{
			args: []string{"10"},
			want: "48",
		},
		{
			args: []string{"283"},
			want: "18480",
		},
		{
			args:     []string{"47547"},
			want:     "96818198400000",
			estimate: 1,
		},
	})
}

/* Brute force approach:
Right triangle generator uses the formula for coprime n:
// https://en.wikipedia.org/wiki/Pythagorean_triple
// a = m^2 - n^2
// b = 2mn
// c = m^2 + n^2
// L = 2m^2 + 2mn
Using this info, for each integer k, we evaluate how many ways k can satisfy
the equation for 'a' or 'b'

func brute176(k int, primes *generator.Generator[int]) int {
	if v, ok := getvCache[k]; ok {
		return v
	}

	cnt := 0

	// (m + n)(m - n)
	for _, f := range generator.Factors(k, primes) {
		d := k / f
		if d <= f {
			continue
		}

		if (d-f)%2 != 0 {
			continue
		}

		n := (d - f) / 2
		m := d - n

		a := 2 * m * n
		b := m*m - n*n
		if !generator.Coprimes(a, b, primes) {
			cnt++
		}
		continue

		if generator.Coprimes(m, n, primes) {
			continue
		}

		cnt++
	}

	// 2mn
	if k%2 == 0 {
		for _, m := range generator.Factors(k/2, primes) {
			n := (k / 2) / m
			if n >= m {
				continue
			}
			a := 2 * m * n
			b := m*m - n*n
			if !generator.Coprimes(a, b, primes) {
				cnt++
			}
		}
	}

	uniqueVCache[k] = cnt

	for _, f := range generator.Factors(k, primes) {
		if f != k {
			cnt += uniqueV(f, primes)
		}
	}

	getvCache[k] = cnt
	return cnt
}

var (
	getvCache = map[int]int{
		0: 0,
		1: 0,
	}
	uniqueVCache = map[int]int{
		0: 0,
		1: 0,
	}
)

func uniqueV(k int, primes *generator.Generator[int]) int {
	if v, ok := uniqueVCache[k]; ok {
		return v
	}
	brute176(k, primes)
	return uniqueVCache[k]
}

*/
