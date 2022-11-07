package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P28() *problem {
	return intInputNode(28, func(o command.Output, n int) {
		start := 3
		sum := 1
		for i := 0; i < (n-1)/2; i++ {
			offset := (i + 1) * 2
			sum += 4*start + 6*offset
			start += offset*4 + 2
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args: []string{"1001"},
			want: "669171001",
		},
		{
			args: []string{"5"},
			want: "101",
		},
	})
}
