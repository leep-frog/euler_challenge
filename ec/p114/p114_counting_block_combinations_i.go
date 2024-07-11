package p114

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P114() *ecmodels.Problem {
	return ecmodels.IntInputNode(114, func(o command.Output, n int) {
		m := map[bool]map[int]int{}
		o.Stdoutln(BlockCombos(false, n, 3, m))
	}, []*ecmodels.Execution{
		{
			Args: []string{"50"},
			Want: "16475640049",
		},
		{
			Args: []string{"7"},
			Want: "17",
		},
	})
}

func BlockCombos(endBlock bool, rem int, minLen int, m map[bool]map[int]int) int {
	if rem < 0 {
		return 0
	}
	if rem == 0 {
		return 1
	}
	if m1, ok := m[endBlock]; ok {
		if v, ok2 := m1[rem]; ok2 {
			return v
		}
	}

	// Add empty block
	v := BlockCombos(false, rem-1, minLen, m)
	if !endBlock {
		// Add blocks of every length
		for j := minLen; j <= rem; j++ {
			v += BlockCombos(true, rem-j, minLen, m)
		}
	}
	if m[endBlock] == nil {
		m[endBlock] = map[int]int{}
	}
	m[endBlock][rem] = v
	return v
}
