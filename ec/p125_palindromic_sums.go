package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P125() *problem {
	return intInputNode(125, func(o command.Output, n int) {
		set := map[int]bool{}
		for i := 1; i*i <= n; i++ {
			sum := i * i
			for j := i + 1; sum+j*j <= n; j++ {
				sum += j * j
				if maths.Palindrome(sum) {
					// Some palindromes can be written in different ways
					set[sum] = true
				}
			}
		}
		var total int
		for k := range set {
			total += k
		}
		o.Stdoutln(total)
	}, []*execution{
		{
			args: []string{"100000000"},
			want: "2906969179",
		},
		{
			args: []string{"1000"},
			want: "4164",
		},
	})
}
