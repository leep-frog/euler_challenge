package eulerchallenge

import (
	"fmt"
	"time"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func cnt1a(n int, primes *generator.Generator[int]) int {
	var sum int
	for iter, prime := primes.Start(0); 2*maths.Pow(prime, 3) <= n; prime = iter.Next() {
		fmt.Println("1a", prime, time.Now())
		sum += generator.PrimePi(n/maths.Pow(prime, 3), primes)
		// Can also check iter
		if maths.Pow(prime, 4) <= n {
			sum--
		}
	}
	return sum
}

// a^3 * b^1
func cnt1(n int, primes *generator.Generator[int]) int {
	var sum int

	fmt.Println("CNT1.1")
	pi := -1
	mx := 1_000
	for ; maths.Pow(primes.Nth(pi+1), 3) < n; pi++ {
		if pi >= mx {
			mx += 1_000
			fmt.Println("PI", pi, primes.Nth(pi))
		}
	}
	fmt.Println("CNT1.2")

	qi := -1

	// Now count down
	for ; pi >= 0; pi-- {
		// Increase qi
		ppp := maths.Pow(primes.Nth(pi), 3)
		fmt.Println("PII", pi, ppp, n/ppp)
		for ; ppp*primes.Nth(qi+1) < n; qi++ {
		}

		// TODO: make prime generator that is bounded by a certain maximum

		if qi < 0 {
			continue
		}
		sum += qi + 1
		if qi >= pi {
			sum--
		}
	}

	return sum
}

// a*b*c
func cnt2a(n int, primes *generator.Generator[int]) int {
	var sum int

	// Do every pair of primes
	for pi := 0; primes.Nth(pi)*primes.Nth(pi+1)*primes.Nth(pi+2) <= n; pi++ {
		fmt.Println("2b", pi, primes.Nth(pi))
		pRem := n / primes.Nth(pi)
		for qi := pi + 1; primes.Nth(qi)*primes.Nth(qi+1) <= pRem; qi++ {
			pqRem := pRem / primes.Nth(qi)
			if pi <= 5 {
				fmt.Println("2bb", primes.Nth(pi), primes.Nth(qi), pqRem, time.Now())
			}
			primeCnt := generator.PrimePi(pqRem, primes)
			// Only count numbers greater than qi
			sum += maths.Max(0, primeCnt-qi-1)
		}
	}

	return sum
}

func cnt2(n int, primes *generator.Generator[int]) int {
	var sum int
	// pi is the smallest
	for pi := 0; primes.Nth(pi)*primes.Nth(pi+1)*primes.Nth(pi+2) < n; pi++ {

		rem := n / primes.Nth(pi)
		qi := pi
		for ; primes.Nth(qi+1)*primes.Nth(qi+2) < rem; qi++ {
		}
		if qi == pi {
			continue
		}

		ri := qi + 1

		for ; qi > pi; qi-- {
			// Increase ri
			for ; primes.Nth(qi)*primes.Nth(ri+1) <= rem; ri++ {
			}
			sum += ri - qi
			// fmt.Println("TWO", pi, qi, ri, primes.Nth(pi), primes.Nth(qi), primes.Nth(ri))
		}
	}
	return sum
}

// a^7
func cnt3(n int, primes *generator.Generator[int]) int {
	var sum int
	for g, p := primes.Start(0); maths.Pow(p, 7) <= n; p = g.Next() {
		sum++
	}
	return sum
}

func P501() *problem {
	return intInputNode(501, func(o command.Output, n int) {
		// All numbers of the pattern (p^3 * q) or (p * q * r) where all are prime
		fmt.Println("START")

		/*for i := 0; i < n; i++ {
			if i%1_000_000_000 == 0 {
				fmt.Println(i)
			}
		}*/

		// primes := generator.PrimesUpTo(n / 5)
		// primes := generator.Primes().Generator
		primes := generator.FinalPrimes(maths.Sqrt(n))

		/*for i := 1; i <= n; i++ {
			// fmt.Println("C1")
			a := cnt1(i, 3, primes, -1)
			// fmt.Println("C2")
			b := cnt2(i, primes)
			// fmt.Println("C3")
			c := cnt3(i, primes)
			o.Stdoutln(i, a, b, c, a+b+c)
		}*/

		/*fmt.Println("C1", time.Now())
		a := cnt1(n, primes)
		fmt.Println("C2", time.Now())
		b := cnt2(n, primes)
		fmt.Println("C3", time.Now())
		c := cnt3(n, primes)
		o.Stdoutln(n, a, cnt1a(n, primes), b, cnt2a(n, primes), c, a+b+c)*/

		fmt.Println("BB", time.Now())
		bb := cnt2a(n, primes)
		fmt.Println("AA", time.Now())
		aa := cnt1a(n, primes)
		fmt.Println("CC", time.Now())
		cc := cnt3(n, primes)
		o.Stdoutln(n, aa, bb, cc, aa+bb+cc)

	}, []*execution{
		{
			args:     []string{"1_000_000_000_000"},
			want:     "197912312715",
			estimate: 75 * 60,
		},
	})
}
