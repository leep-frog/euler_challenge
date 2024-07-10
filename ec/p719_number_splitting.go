package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P719() *problem {
	return intInputNode(719, func(o command.Output, n int) {
		upTo := maths.Pow(10, n)
		var sum int
		for k := 2; k*k <= upTo; k++ {
			digits := maths.Digits(k * k)
			var remNum []int
			for i := 0; i < len(digits); i++ {
				remNum = append(remNum, join(digits[i:]))
			}
			if dp719(digits, remNum, 0, 0, 0, k) {
				sum += k * k
			}
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args: []string{"4"},
			want: "41333",
		},
		{
			args:     []string{"12"},
			want:     "128088830547982",
			estimate: 5,
		},
	})
}

func dp719(digits, remNum []int, idx, count, sum, target int) bool {
	if idx == len(digits) {
		return sum == target && count > 1
	}

	if sum > target {
		return false
	}

	if sum+remNum[idx] < target {
		return false
	}

	for i := idx + 1; i <= len(digits); i++ {
		if dp719(digits, remNum, i, count+1, sum+join(digits[idx:i]), target) {
			return true
		}
	}

	return false
}

func join(is []int) int {
	var sum int
	prod := 1
	for i := len(is) - 1; i >= 0; i-- {
		sum += prod * is[i]
		prod *= 10
	}
	return sum
}
