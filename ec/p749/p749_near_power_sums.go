package p749

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/slices"
)

func P749() *ecmodels.Problem {
	return ecmodels.IntInputNode(749, func(o command.Output, n int) {
		o.Stdoutln(check(n, 0, nil))
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
			Estimate: 45,
		},
	})
}

func check(length, min int, digits []int) int {
	if length == 0 {
		return checkDigits(digits)
	}

	var cnt int
	for i := min; i <= 9; i++ {
		cnt += check(length-1, i, append(digits, i))
	}
	return cnt
}

func checkDigits(digits []int) int {
	if bread.Sum(digits) == 0 {
		return 0
	}

	digitMap := toDigitMap(digits)

	var sum int
	prevPowSum := -1
	for pow := 1; ; pow++ {
		var powSum int
		for k, v := range digitMap {
			powSum += v * maths.Pow(k, pow)
		}

		if len(maths.Digits(powSum-1)) > len(digits) || powSum == prevPowSum {
			break
		}
		prevPowSum = powSum

		for _, offset := range []int{-1, 1} {
			k := powSum + offset
			kDigitMap := toDigitMap(maths.Digits(k))
			if slices.Equal(digitMap, kDigitMap) {
				sum += k
			}
		}

	}
	return sum
}

func toDigitMap(digits []int) []int {
	digitMap := make([]int, 10, 10)
	for _, d := range digits {
		digitMap[d]++
	}
	digitMap[0] = 0
	return digitMap
}
