package p650

import (
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

		psN := p.PrimeFactoredNumberFast(1)

		sum := 1

		factorsOfLowerN := p.PrimeFactoredNumberFast(1)
		for k := 2; k <= n; k++ {

			fdN := p.PrimeFactoredNumberFast(k).Pow(k - 1)

			factorsOfLowerN = factorsOfLowerN.Times(p.PrimeFactoredNumberFast(k - 1))

			psN = psN.Div(factorsOfLowerN).Times(fdN)

			sum = (sum + psN.NumFactors(p, mod)) % mod
		}
		o.Stdoutln(sum)
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
			Args:     []string{"20000"},
			Want:     "538319652",
			Estimate: 30,
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

func brute(n int) int {
	k := maths.One()
	for i := 0; i <= n; i++ {
		k = k.Times(maths.Choose(n, i))
	}

	p := generator.Primes()
	return bread.Sum(bread.Copy(p.Factors(k.ToInt())))
}
