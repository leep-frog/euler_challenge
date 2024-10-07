package p810

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/binary"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P810() *ecmodels.Problem {
	return ecmodels.IntInputNode(810, func(o command.Output, n int) {
		o.Stdoutln(xorSieve(n))
	}, []*ecmodels.Execution{
		{
			Args: []string{"10"},
			Want: "41",
		},
		{
			Args: []string{"10000"},
			Want: "151369",
		},
		{
			Args:     []string{"5000000"},
			Want:     "124136381",
			Estimate: 3.5,
		},
	})
}

func xorSieve(n int) int {

	sizePow := binary.BinaryFromInt(n).Size() + 4
	sieve := make([]bool, maths.Pow(2, sizePow))
	xorPrimes := []int{2}

	sieve[0] = true
	sieve[1] = true
	sieve[2] = false
	sieveIdx := 3

	for len(xorPrimes) < n {
		explore(len(xorPrimes)-1, xorPrimes[len(xorPrimes)-1], xorPrimes, sieve)

		for ; sieve[sieveIdx]; sieveIdx++ {
		}

		xorPrimes = append(xorPrimes, sieveIdx)
	}
	return xorPrimes[len(xorPrimes)-1]
}

func explore(maxIdx, k int, xorPrimes []int, sieve []bool) {
	sieve[k] = true

	for i := 0; i <= maxIdx; i++ {
		// Note that xorPrimes[i] will always be less than k
		// (which is important for xorMult performance reasons).
		v := xorMult(k, xorPrimes[i])
		if v >= len(sieve) {
			break
		}
		explore(i, v, xorPrimes, sieve)
	}
}

// xorMult runs the XOR multiplication logic as defined in the problem.
// For performance reasons, it's best if bi < ai.
func xorMult(ai, bi int) int {
	var r int
	for bi > 0 {
		if bi&1 == 1 {
			r ^= ai
		}
		ai <<= 1
		bi >>= 1
	}
	return r
}

/*
Originally tried to divide by lower numbers. Ultimately too slow, but it was
fun and challenging to implement henve why I left the artifact here.

func isXorDivisble(zi, xi int) (int, bool) {

	z, x := binary.BinaryFromInt(zi).ToDigits(), binary.BinaryFromInt(xi).ToDigits()

	var y []bool
	sums := make([]int, len(x)-1)
	for i := 0; i < len(z)-len(x)+1; i++ {
		sum := sums[0]
		sums = append(sums, 0)[1:]

		if z[i] {
			sum++
		}

		nextY := sum%2 == 1

		if nextY {
			for xi, xv := range x[1:] {
				if xv {
					sums[xi]++
				}
			}
		}

		y = append(y, nextY)
	}

	for i, sum := range sums {
		expectedValue := z[len(z)-len(x)+1+i]
		actualValue := sum%2 == 1
		if expectedValue != actualValue {
			return 0, false
		}
	}

	return binary.FromDigits(y).ToInt(), true
}
*/
