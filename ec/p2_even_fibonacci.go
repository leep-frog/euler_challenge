package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P2() *problem {
	return intInputNode(2, func(o command.Output, n int) {
		var sum int
		for iter, i := generator.Fibonaccis().Start(0); i < n; i = iter.Next() {
			if i%2 == 0 {
				sum += i
			}
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args: []string{"4000000"},
			want: "4613732",
		},
	})
}
