package eulerchallenge

import (
	"fmt"
	"sort"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/maths"
)

func P122() *problem {
	return noInputNode(122, func(o command.Output) {
		var sum int
		for k := 1; k <= 200; k++ {
			_, dist := bfs.ShortestPath([]*node122{{map[int]bool{1: true}, k}}, 0)
			sum += dist
			fmt.Println(k, dist)
		}
		fmt.Println(sum)
	})
}

type node122 struct {
	options map[int]bool
	n       int
}

func (n *node122) Code(*bfs.Context[int, *node122]) string {
	var opts []int
	for o := range n.options {
		opts = append(opts, o)
	}
	sort.Ints(opts)
	c := fmt.Sprintf("%v", opts)
	//fmt.Println(c)
	return c
}

func (n *node122) Done(*bfs.Context[int, *node122]) bool {
	return n.options[n.n]
}

// TODO: prepop and post pop for bfs.AnyPath
func (n *node122) AdjacentStates(*bfs.Context[int, *node122]) []*node122 {
	var r []*node122
	for oi := range n.options {
		for oj := range n.options {
			if oj > oi {
				continue
			}
			if n.options[oi+oj] {
				continue
			}
			cp := maths.CopyMap(n.options)
			cp[oi+oj] = true
			r = append(r, &node122{cp, n.n})
		}
	}
	return r
}
