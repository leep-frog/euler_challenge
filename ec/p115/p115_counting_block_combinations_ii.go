package p115

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/ec/p114"
)

func P115() *ecmodels.Problem {
	return ecmodels.IntInputNode(115, func(o command.Output, m int) {
		n := 1
		cache := map[bool]map[int]int{}
		for ; p114.BlockCombos(false, n, m, cache) < 1_000_000; n++ {
		}
		o.Stdoutln(n)
	}, []*ecmodels.Execution{
		{
			Args: []string{"50"},
			Want: "168",
		},
		{
			Args: []string{"10"},
			Want: "57",
		},
		{
			Args: []string{"3"},
			Want: "30",
		},
	})
}
