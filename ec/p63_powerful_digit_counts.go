package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P63() *problem {
	return noInputNode(63, func(o command.Output) {
		var count int
		for pow := 1; pow < 50; pow++ {
			n := 1
			for ; len(maths.BigPow(n, pow).String()) < pow; n++ {
			}
			for ; len(maths.BigPow(n, pow).String()) == pow; n++ {
				count++
			}
		}
		o.Stdoutln(count)
	})
}
