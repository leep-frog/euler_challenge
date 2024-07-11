package p156

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

// numberOfDigitAppearances returns the number of times the digit d has appeared
// in numbers up to and including k.
func numberOfDigitAppearances(d, k int) int {
	cnt := 0
	for start := 1; start <= k; start *= 10 {
		left := k / (10 * start)
		number := (k / start) % 10
		right := k % start
		cnt += left * start
		if number == d {
			cnt += right + 1
		} else if number > d {
			cnt += start
		}
	}
	return cnt
}

// calculateDigitCountsEqualNumber returns the sum of numbers where the number
// of digit appearances equals the number itself.
func calculateDigitCountsEqualNumber(d, min, minDs, max, maxDs int) int {
	if maxDs < min {
		return 0
	}
	if minDs > max {
		return 0
	}
	if min > max {
		return 0
	}

	if min == max {
		if min == minDs {
			return min
		}
		return 0
	}

	mid := (max + min) / 2
	return (calculateDigitCountsEqualNumber(d, mid+1, numberOfDigitAppearances(d, mid+1), max, maxDs) +
		calculateDigitCountsEqualNumber(d, min, minDs, mid, numberOfDigitAppearances(d, mid)))
}

func P156() *ecmodels.Problem {
	return ecmodels.IntInputNode(156, func(o command.Output, n int) {

		sum := 0
		// Just guessed a really high number. Solution is log(mv) so not a big deal of how big this is.
		mv := 1_000_000_000_000_000
		for i := 1; i <= n; i++ {
			sum += calculateDigitCountsEqualNumber(i, 0, 0, mv, numberOfDigitAppearances(i, mv))
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "22786974071",
		},
		{
			Args: []string{"9"},
			Want: "21295121502550",
		},
	})
}
