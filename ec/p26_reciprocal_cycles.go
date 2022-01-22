package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func cycleLen(num int) int {
	rem := 1
	remMap := map[int]int{}

	for pos := 0; ; pos++ {
		rem = (rem % num) * 10
		if rem == 0 {
			return 0
		}
		if v, ok := remMap[rem]; ok {
			return pos - v
		}
		remMap[rem] = pos
	}
}

func P26() *problem {
	return intInputNode(26, func(o command.Output, n int) {
		var max, maxI int
		for i := 1; i < n; i++ {
			if v := cycleLen(i); v > max {
				max = v
				maxI = i
			}
		}
		o.Stdoutln(maxI)
	})
}
