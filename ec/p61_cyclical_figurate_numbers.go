package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
)

func P61() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=61"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)
			o.Stdoutln(n)

			/*generators := map[int]*generator.Generator[int]{}
			startMap := map[int]int{}
			for i := 0; i < n; i++ {
				shape := i + 3
				g := generator.ShapeNumberGenerator(shape)
				generators[shape] = g
				start := 0
				for ; g.Nth(start) < 1000; start++ {}
				startMap[shape] = start
			}

			for i := 0; generators[3].Nth(i) < 10_000; i++ {

			}*/
		}),
	)
}

type cycFigNum struct {
	n int
	remainingShapes map[int]*generator.Generator[int]
}

func (cfn *cycFigNum) Cycles(that *cycFigNum) bool {
	return cfn.String()[:2] == that.String()[2:]
}

func (cfn *cycFigNum) CyclesInt(that int) bool {
	return cfn.String()[:2] == fmt.Sprintf("%d", that)[2:]
}

func (cfn *cycFigNum) Code(*bfs.Context[int, *cycFigNum]) string {
	return cfn.String()
}

	func (cfn *cycFigNum) String() string {
	return fmt.Sprintf("%d", cfn.n)
}

func (cfn *cycFigNum) Done(ctx *bfs.Context[int, *cycFigNum]) bool {
	var first *cycFigNum
	for cur := ctx.StateValue; cur != nil; cur = cur.Prev() {
		first = cur.State()
	}
	return len(cfn.remainingShapes) == 0 && cfn.Cycles(first)
}

func (cfn *cycFigNum) AdjacentStates(ctx *bfs.Context[map[int]int, *cycFigNum]) []*cycFigNum {
	startMap := ctx.GlobalContext
	var r []*cycFigNum
	for shape, gen := range cfn.remainingShapes {
		for i := startMap[shape]; gen.Nth(i) < 10_000; i++ {
			gn := gen.Nth(i)
			if cfn.CyclesInt(gn) {
				cfn := &cycFigNum{
					gn,
					map[int]*generator.Generator[int]{},
				}
				for k, v := range cfn.remainingShapes {
					if k != shape {
						cfn.remainingShapes[k] = v
					}
				}
				r = append(r, cfn)
			}
		}
	}
	return r
}