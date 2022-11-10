package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/point"
)

func P252() *problem {
	return intInputNode(252, func(o command.Output, n int) {
		o.Stdoutln(n)
	}, []*execution{})
}

type node252 struct {
	points []*point.Point2D[int]
}
