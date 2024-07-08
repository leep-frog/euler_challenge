package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P12() *problem {
	return intInputNode(12, func(o command.Output, n int) {
		ts := generator.Triangulars()
		for idx, i := 0, ts.Nth(0); ; idx, i = idx+1, ts.Nth(idx+1) {
			var count int
			max := i / 2
			for j := 1; j < max; j++ {
				if i%j == 0 {
					max = (i / j) - 1
					if j*j == i {
						count++
					} else {
						count += 2
					}
					if count > n {
						o.Stdoutln(i)
						return
					}
				}
			}
		}
	}, []*execution{
		{
			args: []string{"5"},
			want: "28",
		},
		{
			args:     []string{"500"},
			want:     "76576500",
			estimate: 0.2,
		},
	})
}
