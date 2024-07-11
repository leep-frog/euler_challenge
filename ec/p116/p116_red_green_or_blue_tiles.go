package p116

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P116() *ecmodels.Problem {
	return ecmodels.IntInputNode(116, func(o command.Output, n int) {
		var sum int
		for i := 2; i <= 4; i++ {
			sum += FixedLenBlockCombos(false, n, []int{i}, map[bool]map[int]int{})
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"50"},
			Want: "20492570929",
		},
		{
			Args: []string{"5"},
			Want: "12",
		},
	})
}

func FixedLenBlockCombos(hasOne bool, rem int, lengths []int, m map[bool]map[int]int) int {
	if rem < 0 {
		return 0
	}
	if rem == 0 {
		if hasOne {
			return 1
		}
		return 0
	}

	if m1, ok := m[hasOne]; ok {
		if v, ok2 := m1[rem]; ok2 {
			return v
		}
	}

	// Add empty block
	v := FixedLenBlockCombos(hasOne, rem-1, lengths, m)
	for _, k := range lengths {
		v += FixedLenBlockCombos(true, rem-k, lengths, m)
	}
	if m[hasOne] == nil {
		m[hasOne] = map[int]int{}
	}
	m[hasOne][rem] = v
	return v
}
