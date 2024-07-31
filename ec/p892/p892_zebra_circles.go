package p892

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

const (
	mod = 1234567891
)

func P892() *ecmodels.Problem {
	return ecmodels.NoInputWithExampleNode(892, func(o command.Output, ex bool) {
		if ex {
			o.Stdoutln(solveCircle(100, map[string]map[int]int{}))
			return
		}

		sum := 6
		for i := 4; i <= maths.Pow(10, 7); i++ {
			sum = (sum + d(i)) % mod
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"-x"},
			Want:     "1172122931",
			Estimate: 1.5,
		},
		{
			Want:     "469137427",
			Estimate: 10,
		},
	})
}

/*******************
* Initial approach *
********************/

// After looking at brute force values, I noticed that the even and odd sequences of D(n)
// resulted in modifications of existing OEIS sequences

func d(i int) int {
	if i%2 == 1 {
		return dOdd(i / 2)
	}
	return dEven(i / 2)
}

var (
	// From https://oeis.org/A000108
	// Recurrence: a(n) = 2*(2*n-1)*a(n-1)/(n+1) with a(0) = 1.
	prevDD = 1
)

func dOdd(i int) int {
	// a(n) = a(n-1)*( 4*(4*n^2-1)/(n*(n+2)) )
	n := i - 1

	num := (4 * ((4*n*n - 1) % mod) % mod)
	den := (n * (n + 2)) % mod
	prevDD = (prevDD * num) % mod
	prevDD = (prevDD * inverse(den)) % mod
	return (prevDD * 4 * (2*i - 1)) % mod
}

var (
	// From https://oeis.org/A060150
	// Recurrence: n^2*a(n) -4*(2*n-1)^2*a(n-1)=0
	//  a(n) = 4*(2*n-1)^2*a(n-1) / n^2
	prevDD2 = 2
)

func dEven(i int) int {
	n := i
	num := (((4 * (2*n - 1)) % mod) * (2*n - 1)) % mod
	den := (n * n) % mod
	prevDD2 = (prevDD2 * num) % mod
	prevDD2 = (prevDD2 * inverse(den)) % mod
	return prevDD2
}

func inverse(k int) int {
	return maths.PowMod(k, -1, mod)
}

/*******************
* Initial approach *
********************/

func solveCircle(size int, cache map[string]map[int]int) int {
	m := solveCircleDp(size*2, false, cache)
	var sum int
	for v, cnt := range m {
		sum = (sum + maths.Abs(v*cnt)) % mod
	}
	return sum
}

func solveCircleDp(size int, inWhite bool, cache map[string]map[int]int) map[int]int {
	// Pair up each group
	if size%2 == 1 {
		panic("Size should be even")
	}

	if size == 0 {
		v := 1
		if !inWhite {
			v = -1
		}
		return map[int]int{v: 1}
	}

	code := fmt.Sprintf("%d-%v", size, inWhite)
	if m, ok := cache[code]; ok {
		return m
	}

	m := map[int]int{}
	for i := 1; i < size; i += 2 {

		leftWhite, rightWhite := inWhite, !inWhite
		lefts := solveCircleDp(i-1, leftWhite, cache)
		rights := solveCircleDp(size-i-1, rightWhite, cache)
		for leftValue, leftCnt := range lefts {
			for rightValue, rightCnt := range rights {
				a := (leftCnt * rightCnt) % mod
				m[leftValue+rightValue] = (m[leftValue+rightValue] + a) % mod
			}
		}
	}

	cache[code] = m
	return m
}
