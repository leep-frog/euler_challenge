package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P48() *problem {
	return intInputNode(48, func(o command.Output, n int) {
		res := 0
		largest := 10_000_000_000
		_ = largest
		for i := 1; i <= n; i++ {
			prod := i
			for j := 1; j < i; j++ {
				prod = (prod * i) % largest
			}
			res = (res + prod) % largest
		}
		o.Stdoutln(res % largest)
	})
}
