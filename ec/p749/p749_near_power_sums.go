package p749

import (
	"math"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/slices"
)

func P749() *ecmodels.Problem {
	return ecmodels.IntInputNode(749, func(o command.Output, n int) {

		var parts []int
		for i := 0; i <= 9; i++ {
			parts = append(parts, i)
		}

		c := &combinatorics.Combinatorics[int]{
			Parts:            parts,
			MinLength:        n,
			MaxLength:        n,
			AllowReplacement: true,
			OrderMatters:     false,
		}

		maxValue := maths.Pow(10, n)

		var sum int
		combinatorics.EvaluateCombos(c, func(digits []int) {
			sum += checkDigits(maxValue, toDigitMap(digits))
		})
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "110",
		},
		{
			Args: []string{"6"},
			Want: "2562701",
		},
		{
			Args:     []string{"16"},
			Want:     "13459471903176422",
			Estimate: 6,
		},
	})
}

/* This was used initially, but then realized that we can just use the combinatorics package

func check(n, length, min int, digitMap []int) int {
	if length == 0 {
		return checkDigits(n, digitMap)
	}

	var cnt int
	for i := min; i <= 9; i++ {
		digitMap[i]++
		cnt += check(n, length-1, i, digitMap)
		digitMap[i]--
	}
	return cnt
}*/

func checkDigits(maxValue int, digitMap []int) int {
	pos := bread.Sum(digitMap[1:])
	if pos == 0 {
		return 0
	}

	// Let pos be the number of non-zero values. Then, the largest the powSum can be
	// is pos*(9^pow). This value should be at least 10^(pos).
	// So, solve for pow in the below:
	// pos*9^pow >= 10^pos
	// log(pos) + pow*log(9) >= pos*log(10)
	// pow*log(9) >= pos*log(10) - log(pos)
	// pow >= (pos*log(10) - log(pos)) / log(9)
	startPow := int(math.Ceil(((float64(pos)*math.Log(10) - math.Log(float64(pos))) / math.Log(9))))

	var sum int
	prevPowSum := -1
	for pow := maths.Max(startPow, 1); ; pow++ {
		var powSum int
		for k, v := range digitMap {
			powSum += v * maths.Pow(k, pow)
		}

		if powSum >= maxValue || powSum == prevPowSum {
			break
		}
		prevPowSum = powSum

		for _, offset := range []int{-1, 1} {
			k := powSum + offset
			kDigitMap := toDigitMap(maths.Digits(k))

			// The number of zeroes can be different (e.g. [3 5], [0 0 3 5]), so
			// ignore those in the comparison
			if slices.Equal(digitMap[1:], kDigitMap[1:]) {
				sum += k
			}
		}

	}
	return sum
}

func toDigitMap(digits []int) []int {
	digitMap := make([]int, 10)
	for _, d := range digits {
		digitMap[d]++
	}
	return digitMap
}
