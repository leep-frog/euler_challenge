package eulerchallenge

import (
	"strconv"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bfs"
)

func P122() *problem {
	return noInputNode(122, func(o command.Output) {
		var sum int
		for k := 1; k <= 200; k++ {
			for i := 1; ; i++ {
				path := bfs.ContextualDFS([]*node122{{1}}, &context122{k, i}, bfs.AllowDFSCycles(), bfs.AllowDFSDuplicates())
				if len(path) != 0 {
					sum += len(path) - 1
					break
				}
			}
		}
		o.Stdoutln(sum)
	}, &execution{
		want:     "1582",
		estimate: 5,
	})
}

type node122 struct {
	pow int
}

type context122 struct {
	pow      int
	maxDepth int
}

func (n *node122) String() string {
	return strconv.Itoa(n.pow)
}

func (n *node122) Code(ctx *context122, path bfs.DFSPath[*node122]) string {
	return strconv.Itoa(n.pow)
}

func (n *node122) Done(ctx *context122, path bfs.DFSPath[*node122]) bool {
	return n.pow == ctx.pow
}

func (n *node122) AdjacentStates(ctx *context122, path bfs.DFSPath[*node122]) []*node122 {
	if path.Len() >= ctx.maxDepth {
		return nil
	}
	var r []*node122
	for _, p := range path.Path() {
		if path.Contains(strconv.Itoa(n.pow + p.pow)) {
			continue
		}
		if n.pow+p.pow > ctx.pow {
			continue
		}
		r = append(r, &node122{p.pow + n.pow})
	}
	return r
}
