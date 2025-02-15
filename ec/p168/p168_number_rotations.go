package p168

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func check168(maxSize, factor, rightDigit int) int {
	digits := []int{rightDigit}
	var rem int
	var sum int
	for j := 0; j < maxSize-1; j++ {
		product := (digits[len(digits)-1] * factor) + rem
		digits = append(digits, (product % 10))
		rem = product / 10

		if digits[len(digits)-1] != 0 && rem+digits[len(digits)-1]*factor == digits[0] {
			sum += maths.FromDigits(bread.Reverse(digits[:maths.Min(len(digits), 5)]))
		}
	}
	return sum
}

func P168() *ecmodels.Problem {
	return ecmodels.IntInputNode(168, func(o command.Output, n int) {
		var sum int
		for factor := 1; factor <= 9; factor++ {
			for rightDigit := 1; rightDigit <= 9; rightDigit++ {
				sum += check168(n, factor, rightDigit)
			}
		}

		o.Stdoutln(sum % 100_000)
	}, []*ecmodels.Execution{
		{
			Args: []string{"6"},
			Want: "98331",
		},
		{
			Args: []string{"100"},
			Want: "59206",
		},
	})
}
