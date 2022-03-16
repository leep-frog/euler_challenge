package eulerchallenge

import (
	"fmt"
	
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/maths"
)

func P119() *problem {
	return intInputNode(119, func(o command.Output, n int) {
		ctx := &context119{nil, n}
		bfs.ShortestWeightedPath([]*node119{{2, 2}}, ctx)
		o.Stdoutln(ctx.values[n-1])
	})
}

type node119 struct {
	int_ int
	pow int
}

type context119 struct {
	values []*maths.Int
	n int
}

func (n *node119) Code(*bfs.Context[*context119, *node119]) string {
	return fmt.Sprintf("%v^%d", n.int_, n.pow)
}

func (n *node119) Distance(ctx *bfs.Context[*context119, *node119]) int {
	return len(maths.BigPow(n.int_, n.pow).Digits())
}

func (n *node119) Done(ctx *bfs.Context[*context119, *node119]) bool {
	if maths.SumSys(maths.BigPow(n.int_, n.pow).Digits()...) == n.int_ {
		ctx.GlobalContext.values = append(ctx.GlobalContext.values,  maths.BigPow(n.int_, n.pow))
	}
	return len(ctx.GlobalContext.values) >= ctx.GlobalContext.n 
}

func (n *node119) AdjacentStates(ctx *bfs.Context[*context119, *node119]) []*node119 {
	return []*node119{
		{n.int_+1, n.pow},
		{n.int_, n.pow+1},
	}
}