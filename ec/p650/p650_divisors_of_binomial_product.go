package p650

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

const mod = 1_000_000_007

func P650() *ecmodels.Problem {
	return ecmodels.IntInputNode(650, func(o command.Output, n int) {
		p := generator.Primes()
		_ = p

		// fmt.Println("A", time.Now())
		// for i := 1; i <= n; i++ {
		// 	for _, pf := range p.PrimeFactorsFast(i) {
		// 		_ = pf[0] + pf[1]
		// 	}
		// 	// fmt.Println(i, p.PrimeFactorsFast(i), p.PrimeFactors(i))
		// }

		// fmt.Println("B", time.Now())
		// generator.ClearCaches()
		// for i := 1; i <= n; i++ {
		// 	p.PrimeFactors(i)
		// }
		// fmt.Println("C", time.Now())
		// return

		// for a := 2; a <= n; a++ {
		// 	for b := 1; b <= n; b++ {
		// 		fmt.Println(a, b, harm(a, b), harm2(a, b))
		// 	}
		// }

		// o.Stdoutln(clever2(n))
		fmt.Println(clever2(n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"5"},
			Want: "5736",
		},
		{
			Args: []string{"10"},
			Want: "721034267",
		},
		{
			Args: []string{"100"},
			Want: "332792866",
		},
		{
			Args: []string{"20000"},
			Want: "538319652",
		},
	})
}

// func basic(m int) int {
// 	p := generator.Primes()
// 	_ = p

// 	var sum int
// 	for n := 1; n <= m; n++ {

// 		counts := make([]int, n+1)

// 		for i := 0; i <= n; i++ {
// 			counts[i] = n - 1
// 		}

// 		for i := 1; i < n; i++ {
// 			coef := 2 * (n - i)
// 			counts[i] -= coef
// 		}

// 		ps := map[int]int{}
// 		for i := 1; i <= n; i++ {
// 			for k, v := range p.MutablePrimeFactors(i) {
// 				ps[k] += counts[i] * v
// 				if ps[k] == 0 {
// 					delete(ps, k)
// 				}
// 			}
// 		}

// 		var pfPairs [][]int
// 		for k, v := range ps {
// 			if v != 0 {
// 				pfPairs = append(pfPairs, []int{k, v})
// 			}
// 		}

// 		// fmt.Println(divisorSum(0, pfPairs)+1, brute(n))
// 		// fmt.Println(n, divisorSum(0, pfPairs)+1)
// 		sum = (sum + divisorSum2(0, pfPairs)) % mod
// 		// fmt.Println(n, pfPairs, divisorSum2(0, pfPairs))
// 		fmt.Println("REGULAR", n, sum, pfPairs)
// 	}
// 	// 721034267
// 	return sum
// }

// func clever(n int) int {
// 	p := generator.Primes()
// 	ps := map[int]int{}

// 	sum := 1
// 	for k := 2; k <= n; k++ {
// 		for ky, v := range p.PrimeFactorsFast(k) {
// 			ps[ky] += v * (k - 2)
// 		}

// 		for b := 1; b < k-1; b++ {
// 			for k, v := range p.PrimeFactorsFast(k - b) {
// 				ps[k] -= v
// 			}
// 		}

// 		for k, v := range p.PrimeFactorsFast(k) {
// 			ps[k] += v
// 		}

// 		for k := range maps.Keys(ps) {
// 			if ps[k] == 0 {
// 				delete(ps, k)
// 			}
// 		}

// 		var pfPairs [][]int
// 		for k, v := range ps {
// 			if v != 0 {
// 				pfPairs = append(pfPairs, []int{k, v})
// 			}
// 		}
// 		sum = (sum + divisorSum2(0, pfPairs)) % mod
// 		fmt.Println("CLEVER", k, sum)
// 	}
// 	return sum
// }

func clever2(n int) int {
	p := generator.Primes()
	ps := map[int]int{}

	sum := 1
	prevSum := 1

	factorsOfLower := map[int]int{}
	for k := 2; k <= n; k++ {

		fd := map[int]int{}
		for _, pff := range p.PrimeFactorsFast(k) {
			ky := pff[0]
			v := pff[1]

			fd[ky] = v * (k - 1)
		}

		for _, pff := range p.PrimeFactorsFast(k - 1) {
			k := pff[0]
			v := pff[1]
			factorsOfLower[k] += v
		}

		for k, v := range factorsOfLower {
			fd[k] -= v
		}

		if k%1000 == 0 {
			fmt.Println(k)
		}

		for k, v := range fd {
			old := harm(k, ps[k])
			new := harm(k, ps[k]+v)
			prevSum = (((prevSum * inv(old)) % mod) * new) % mod

			ps[k] += v
			if ps[k] == 0 {
				delete(ps, k)
			}
		}

		sum = (sum + prevSum) % mod
	}
	return sum
}

func brute(n int) int {
	k := maths.One()
	for i := 0; i <= n; i++ {
		k = k.Times(maths.Choose(n, i))
	}

	p := generator.Primes()
	return bread.Sum(bread.Copy(p.Factors(k.ToInt())))
}

var invCache = map[int]int{}

func inv(k int) int {
	if v, ok := invCache[k]; ok {
		return v
	}
	v := maths.PowMod(k, -1, mod)
	invCache[k] = v
	return v
}

// for k, v := range invCache {}

func harm(k, pow int) int {
	p := maths.PowMod(k, pow+1, mod)
	p = (p + mod - 1) % mod

	return (p * inv(k-1)) % mod
}

// var(
// 	pmCache = [][]int{}
// )

// func pm(k, pow int)  int {
// 	for len(pmCache) <= k {
// 		pmCache = append(pmCache, []int{})
// 	}
// 	for len(pmCache[k]) <= pow {
// 		pmCache[k] = append(pmCache[k], maths.PowMod)
// 	}
// }
