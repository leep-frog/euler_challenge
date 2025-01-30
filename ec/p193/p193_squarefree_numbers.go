package p193

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P193() *ecmodels.Problem {
	return ecmodels.IntInputNode(193, func(o command.Output, n int) {
		p := generator.Primes()

		pow := maths.Pow(2, n)

		// fmt.Println("BRU", brute(p, pow))
		// fmt.Println("CALC", calc(p, 0, pow))
		// fmt.Println("CALS", pow-calcSquares(p, pow))
		fmt.Println("OTHER", pow-other(p, pow))
		// fmt.Println(oCache)
		// o.Stdoutln(n)
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "",
		},
		{
			Args: []string{"2"},
			Want: "",
		},
	})
}

var (
	cache = map[string]int{}
)

func calc(p *generator.Prime, nextIdx, rem int) int {

	code := fmt.Sprintf("%d-%d", nextIdx, rem)
	if v, ok := cache[code]; ok {
		// fmt.Println("CACHE HIT")
		return v
	}

	sum := 1

	for i := nextIdx; p.Nth(i) <= rem; i++ {
		sum += calc(p, i+1, rem/p.Nth(i))
	}

	cache[code] = sum
	return sum

}

func calcSquares(p *generator.Prime, n int) int {
	// 4 removes every 4th number
	// 9 removes every 9th number (but re-add every 4th combination)

	// (n / 4)
	// + (n / 9) - (n / 36)
	// + (n / 25) - (n /

	// (n / 2)
	// + (n / 3) - (n / 2,3)
	// + (n / 5) - (n / 5,2) - (n / 5,3) + (n / 5,3,2)
	// + (n / 7) - (n / 7,2) - (n / 7,3) - (n / 7,5) + (n / 7,2,3) + (n / 7,2,5) + (n / 7,3,5) - (n / 7,2,3,5)

	// n = 37
	// (n / 4) = 37/4 = 9
	// + (n / 9) - (n / 36) = 4 - 1
	// + (n / 25) = 1
	// = 9 + 4 - 1 + 1 = 13

	// Base case: [0]: {positive = (n/2), negative = 0}
	// Iterative case: [k]: {positive = (n/k) + negative[k-1]/k, negative = positive[k-1]/k}

	/*positive, negative := (n / (p.Nth(0) * p.Nth(0))), 0
	sum := positive
	for i := 1; p.Nth(i)*p.Nth(i) <= n; i++ {
		sq := p.Nth(i) * p.Nth(i)
		positive, negative = (n+negative)/sq, positive/sq
		fmt.Println(positive, negative)
	}
	fmt.Println(positive, negative)
	return positive - negative*/

	vals := []int{n / (p.Nth(0) * p.Nth(0))}
	sum := vals[0]

	for i := 1; p.Nth(i)*p.Nth(i) <= n; i++ {
		sq := p.Nth(i) * p.Nth(i)
		nextVals := []int{n / sq}
		sum += nextVals[0]
		for _, val := range vals {
			nextVal := -val / sq
			if nextVal != 0 {
				sum += nextVal
				nextVals = append(nextVals, nextVal)
			}
		}
		vals = nextVals
	}

	return sum

}

func brute(p *generator.Prime, n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		squareFree := true
		for _, cnt := range p.PrimeFactors(i) {
			if cnt > 1 {
				squareFree = false
				break
			}
		}

		if squareFree {
			sum++
		}
	}
	return sum
}

var (
	oCache = map[string]int{}
)

func other(p *generator.Prime, n int) int {
	var sum int
	for i := 0; p.Nth(i)*p.Nth(i) <= n; i++ {
		sum += otherRecur(p, n, i)
	}
	return sum
}

func otherRecur(p *generator.Prime, n, pIdx int) int {
	if n == 0 {
		return 0
	}

	sq := p.Nth(pIdx) * p.Nth(pIdx)
	sum := n / sq

	code := fmt.Sprintf("%d-%d", p.Nth(pIdx), sum)
	if v, ok := oCache[code]; ok {
		return v
	}

	for i := 0; i < pIdx; i++ {
		v := otherRecur(p, n/sq, i)
		if v == 0 {
			break
		}
		sum -= v
	}

	oCache[code] = sum
	return sum
}
