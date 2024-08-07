package p820

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P820() *ecmodels.Problem {
	return ecmodels.IntInputNode(820, func(o command.Output, n int) {

		// seens := make([]int, n, n)

		var sum int
		for i := 1; i <= n; i++ {
			// sum += nthDigit(i, n, seens[:0])
			sum += clever(i, n)
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"7"},
			Want: "10",
		},
		{
			Args: []string{"100"},
			Want: "418",
		},
		{
			Args:     []string{"10000000"},
			Want:     "44967734",
			Estimate: 20,
		},
	})
}

func clever(k, n int) int {

	// d_n(k) = [ floor(10^n) / k ] mod 10

	// Not sure why (got equality online), but
	// d_n(k) = floor[ (10^n mod 10k) / k ]
	// return maths.PowMod(10, n, 10*k) / k

	// The above can also be simplified to
	// d_n(k) = floor[ 10 * (10^(n-1) mod k) / k ]
	return 10 * maths.PowMod(10, n-1, k) / k
}

func nthDigit(k, n int, seens []int) int {
	num := 1

	if len(seens) > 0 {
		panic("seens should be empty (but have large capacity)")
	}

	seen := map[int]int{}
	var nextDigit int
	for i := 0; i <= n; i++ {
		nextDigit = num / k
		num = (num % k) * 10
		if num == 0 {
			return 0
		}

		seens = append(seens, nextDigit)

		if v, ok := seen[num]; ok {
			patternLen := i - v
			patternOffset := (n - i) % patternLen

			if patternOffset == 0 {
				patternOffset = patternLen
			}

			seensIdx := (len(seens) - 1) - patternLen + patternOffset

			return seens[seensIdx]
		} else {
			seen[num] = i
		}

	}

	return nextDigit
}

func nthDigitBrute(k, n int) int {
	num := 1

	var nextDigit int
	for i := 0; i <= n; i++ {
		nextDigit = num / k
		num = (num % k) * 10
		if num == 0 {
			return 0
		}
	}
	return nextDigit
}
