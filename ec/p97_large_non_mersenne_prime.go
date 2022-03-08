package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P97() *problem {
	return noInputNode(97, func(o command.Output) {
		v := int64(1)
		for i := 0; i < 7830457; i++ {
			v *= 2
			v = v % 1_000_000_000_000_000_000
		}
		v = v % 1_000_000_000_000
		v *= 28433
		v += 1
		vStr := fmt.Sprintf("%d", v)
		k := maths.Max(len(vStr)-10, 0)
		o.Stdoutln(vStr[k:])
	})
}
