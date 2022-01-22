package eulerchallenge

import (
	"github.com/leep-frog/command"
)

const (
	N = "N"
)

func P1() *problem {
	return intInputNode(1, func(o command.Output, n int) {
		var sum int
		for i := 1; i < n; i++ {
			if i%5 == 0 || i%3 == 0 {
				sum += i
			}
		}
		o.Stdoutln(sum)
	})
}
