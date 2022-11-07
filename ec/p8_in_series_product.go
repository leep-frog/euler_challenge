package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
)

func P8() *problem {
	return intInputNode(8, func(o command.Output, n int) {
		s := parse.ReadFileInput("p8.txt")
		var is []int
		for i := 0; i < len(s); i++ {
			is = append(is, parse.Atoi(s[i:i+1]))
		}

		var max int

		for i := n; i < len(s); i++ {
			product := 1
			for j := i - n; j < i; j++ {
				product *= is[j]
			}
			if product > max {
				max = product
			}
		}

		o.Stdoutln(max)
	}, []*execution{
		{
			args: []string{"4"},
			want: "5832",
		},
		{
			args: []string{"13"},
			want: "23514624000",
		},
	})
}
