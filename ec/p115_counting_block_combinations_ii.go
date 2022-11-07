package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P115() *problem {
	return intInputNode(115, func(o command.Output, m int) {
		n := 1
		cache := map[bool]map[int]int{}
		for ; blockCombos(false, n, m, cache) < 1_000_000; n++ {
		}
		o.Stdoutln(n)
	}, []*execution{
		{
			args: []string{"50"},
			want: "168",
		},
		{
			args: []string{"10"},
			want: "57",
		},
		{
			args: []string{"3"},
			want: "30",
		},
	})
}
