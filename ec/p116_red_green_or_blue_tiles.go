package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P116() *problem {
	return intInputNode(116, func(o command.Output, n int) {
		var sum int
		for i := 2; i <= 4; i++ {
			sum += fixedLenBlockCombos(false, n, []int{i}, map[bool]map[int]int{})
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args: []string{"50"},
			want: "20492570929",
		},
		{
			args: []string{"5"},
			want: "12",
		},
	})
}

func fixedLenBlockCombos(hasOne bool, rem int, lengths []int, m map[bool]map[int]int) int {
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
	v := fixedLenBlockCombos(hasOne, rem-1, lengths, m)
	for _, k := range lengths {
		v += fixedLenBlockCombos(true, rem-k, lengths, m)
	}
	if m[hasOne] == nil {
		m[hasOne] = map[int]int{}
	}
	m[hasOne][rem] = v
	return v
}
