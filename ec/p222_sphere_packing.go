package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P222() *problem {
	return intInputNode(222, func(o command.Output, n int) {
		// Triangle is
		o.Stdoutln(n)
	}, []*execution{
		{
			args: []string{"1"},
			want: "0",
			skip: "TODO",
		},
	})
}
