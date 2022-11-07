package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P25() *problem {
	return intInputNode(25, func(o command.Output, n int) {
		for g, i := generator.BigFibonaccis(), 1; ; i++ {
			if len(g.Nth(i).String()) >= n {
				o.Stdoutln(i + 1)
				return
			}
		}
	}, []*execution{
		{
			args:     []string{"1000"},
			want:     "4782",
			estimate: 0.2,
		},
		{
			args: []string{"2"},
			want: "7",
		},
		{
			args: []string{"1"},
			want: "2",
		},
	})
}
