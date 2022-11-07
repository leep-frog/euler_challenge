package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P74() *problem {
	return intInputNode(74, func(o command.Output, n int) {
		chainLength := map[int]int{}
		for i := 1; i < n; i++ {
			curChain := map[int]bool{
				i: true,
			}
			k := factorialDigitSum(i)
			for ; !curChain[k] && chainLength[k] == 0; k = factorialDigitSum(k) {
				curChain[k] = true
			}
			if curChain[k] {
				chainLength[i] = len(curChain)
			} else {
				chainLength[i] = len(curChain) + chainLength[k]
			}
		}

		count := 0
		for _, v := range chainLength {
			if v == 60 {
				count++
			}
		}
		o.Stdoutln(count)
	}, []*execution{
		{
			args:     []string{"1000000"},
			want:     "402",
			estimate: 1,
		},
	})
}

func factorialDigitSum(n int) int {
	var sum int
	for _, d := range maths.Digits(n) {
		sum += maths.FactorialI(d)
	}
	return sum
}
