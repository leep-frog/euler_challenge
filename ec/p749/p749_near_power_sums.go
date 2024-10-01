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
		o.Stdoutln(check(n, n, 0, make([]int, 10)))
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
}

func checkDigits(n int, digitMap []int) int {
	if bread.Sum(digitMap[1:]) == 0 {
		return 0
	}

	old := digitMap[0]
	digitMap[0] = 0

	var sum int
	prevPowSum := -1
	for pow := 1; ; pow++ {
		var powSum int
		for k, v := range digitMap {
			powSum += v * maths.Pow(k, pow)
		}

		if len(maths.Digits(powSum-1)) > n || powSum == prevPowSum {
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
	digitMap[0] = old
	return sum
}

func toDigitMap(digits []int) []int {
	digitMap := make([]int, 10)
	for _, d := range digits {
		digitMap[d]++
	}
	digitMap[0] = 0
	return digitMap
}
