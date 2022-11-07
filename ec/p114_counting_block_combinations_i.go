package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P114() *problem {
	return intInputNode(114, func(o command.Output, n int) {
		m := map[bool]map[int]int{}
		o.Stdoutln(blockCombos(false, n, 3, m))
	}, []*execution{
		{
			args: []string{"50"},
			want: "16475640049",
		},
		{
			args: []string{"7"},
			want: "17",
		},
	})
}

func blockCombos(endBlock bool, rem int, minLen int, m map[bool]map[int]int) int {
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
	v := blockCombos(false, rem-1, minLen, m)
	if !endBlock {
		// Add blocks of every length
		for j := minLen; j <= rem; j++ {
			v += blockCombos(true, rem-j, minLen, m)
		}
	}
	if m[endBlock] == nil {
		m[endBlock] = map[int]int{}
	}
	m[endBlock][rem] = v
	return v
}
