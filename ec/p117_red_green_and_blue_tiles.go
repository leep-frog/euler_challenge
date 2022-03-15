package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P117() *problem {
	return intInputNode(117, func(o command.Output, n int) {
		o.Stdoutln(fixedLenBlockCombos(true, n, []int{2, 3, 4}, map[bool]map[int]int{}))
	})
}
