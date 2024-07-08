package eulerchallenge

import (
	"github.com/leep-frog/command/command"
)

func P117() *problem {
	return intInputNode(117, func(o command.Output, n int) {
		o.Stdoutln(fixedLenBlockCombos(true, n, []int{2, 3, 4}, map[bool]map[int]int{}))
	}, []*execution{
		{
			args: []string{"50"},
			want: "100808458960497",
		},
		{
			args: []string{"5"},
			want: "15",
		},
	})
}
