package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

// a^3 * b^1
func cnt1(n, exp int, primes *generator.Generator[int], min int) int {
	var sum int

	// fmt.Println("CNT1.1")
	pi := min
	mx := 1_000
	for ; maths.Pow(primes.Nth(pi+1), exp) < n; pi++ {
		if pi >= mx {
			mx += 1_000
			// fmt.Println("PI", pi, primes.Nth(pi))
		}
	}
	// fmt.Println("CNT1.2")

	qi := -1

	// Now count down
	for ; pi > min; pi-- {
		// Increase qi
		ppp := maths.Pow(primes.Nth(pi), exp)
		// fmt.Println("PII", pi, ppp, n/ppp)
		for ; ppp*primes.Nth(qi+1) < n; qi++ {
		}

		// TODO: make prime generator that is bounded by a certain maximum

		if qi <= min {
			continue
		}
		sum += qi - min
		if qi >= pi {
			sum--
		}
	}

	return sum
}

// a*b*c
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

		for i := 0; i < n; i++ {
			if i%1_000_000_000 == 0 {
				fmt.Println(i)
			}
		}

		// primes := generator.PrimesUpTo(n / 5)
		primes := generator.Primes().Generator

		/*for i := 1; i <= n; i++ {
			// fmt.Println("C1")
			a := cnt1(i, 3, primes, -1)
			// fmt.Println("C2")
			b := cnt2(i, primes)
			// fmt.Println("C3")
			c := cnt3(i, primes)
			o.Stdoutln(i, a, b, c, a+b+c)
		}*/

		a := cnt1(n, 3, primes, -1)
		// fmt.Println("C2")
		b := cnt2(n, primes)
		// fmt.Println("C3")
		c := cnt3(n, primes)
		o.Stdoutln(n, a, b, c, a+b+c)

	}, []*execution{
		{
			args: []string{"1_000_000_000_000"},
			want: "0",
		},
	})
}
