package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P61() *problem {
	return intInputNode(61, func(o command.Output, n int) {
		generators := map[int]*generator.Generator[int]{}
		startMap := map[int]int{}
		for i := 1; i < n; i++ {
			shape := i + 3
			g := generator.ShapeNumberGenerator(shape)
			generators[shape] = g
			start := 0
			for ; g.Nth(start) < 1000; start++ {
			}
			startMap[shape] = start
		}

		var initStates []*cycFigNum
		triG := generator.ShapeNumberGenerator(3)
		for i := 0; triG.Nth(i) < 10_000; i++ {
			if triG.Nth(i) < 1000 {
				continue
			}
			initStates = append(initStates, &cycFigNum{
				triG.Nth(i),
				maths.CopyMap(generators),
			})
		}
		path, _ := bfs.ContextualShortestPathWithPath[bfs.Int](initStates, startMap, bfs.CheckDuplicates())
		o.Stdoutln(maths.SumType(path))
	}, []*execution{
		{
			args: []string{"6"},
			want: "28684",
		},
		{
			args: []string{"3"},
			want: "19291",
		},
	})
}

type cycFigNum struct {
	n               int
	remainingShapes map[int]*generator.Generator[int]
}

func (cfn *cycFigNum) ToInt() int {
	return cfn.n
}

func (cfn *cycFigNum) Distance(m map[int]int, path bfs.Path[*cycFigNum]) bfs.Int {
	return bfs.Int(-path.Len())
}

func (cfn *cycFigNum) Cycles(that *cycFigNum) bool {
	return cfn.String()[2:] == that.String()[:2]
}

func (cfn *cycFigNum) CyclesInt(that int) bool {
	return cfn.String()[2:] == fmt.Sprintf("%d", that)[:2]
}

func (cfn *cycFigNum) Code(map[int]int, bfs.Path[*cycFigNum]) string {
	return cfn.String()
}

func (cfn *cycFigNum) String() string {
	return fmt.Sprintf("%d", cfn.n)
}

func (cfn *cycFigNum) Done(m map[int]int, path bfs.Path[*cycFigNum]) bool {
	if len(cfn.remainingShapes) > 0 {
		return false
	}
	return cfn.Cycles(path.Fetch()[0])
}

func (cfn *cycFigNum) AdjacentStates(startMap map[int]int, path bfs.Path[*cycFigNum]) []*cycFigNum {
	var r []*cycFigNum
	for shape, gen := range cfn.remainingShapes {
		for i := startMap[shape]; gen.Nth(i) < 10_000; i++ {
			gn := gen.Nth(i)
			if cfn.CyclesInt(gn) {
				r = append(r, &cycFigNum{
					gn,
					maths.CopyMap(cfn.remainingShapes, shape),
				})
			}
		}
	}
	return r
}
