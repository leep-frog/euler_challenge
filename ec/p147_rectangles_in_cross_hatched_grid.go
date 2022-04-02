package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P147() *problem {
	return intInputNode(147, func(o command.Output, n int) {
		o.Stdoutln(n)
	})
}
