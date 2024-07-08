package eulerchallenge

import (
	"github.com/leep-frog/command/command"
)

func P14() *problem {
	return intInputNode(14, func(o command.Output, n int) {
		found := map[int]int{}
		var max, maxI int
		for i := 2; i < n; i++ {
			count := 1
			for j := i; j != 1; {
				if j%2 == 0 {
					j /= 2
				} else {
					j = 3*j + 1
				}

				if v, ok := found[j]; ok {
					count += v
					break
				}
				count++
			}
			found[i] = count
			if count > max {
				max = count
				maxI = i
			}
		}
		o.Stdoutln(maxI)
	}, []*execution{
		{
			args:     []string{"1000000"},
			want:     "837799",
			estimate: 0.4,
		},
	})
}
