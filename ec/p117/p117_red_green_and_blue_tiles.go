package p117

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/ec/p116"
)

func P117() *ecmodels.Problem {
	return ecmodels.IntInputNode(117, func(o command.Output, n int) {
		o.Stdoutln(p116.FixedLenBlockCombos(true, n, []int{2, 3, 4}, map[bool]map[int]int{}))
	}, []*ecmodels.Execution{
		{
			Args: []string{"50"},
			Want: "100808458960497",
		},
		{
			Args: []string{"5"},
			Want: "15",
		},
	})
}
